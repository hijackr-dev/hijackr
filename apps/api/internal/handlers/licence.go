package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ValidateLicenceRequest struct {
	LicenceKey  string `json:"licence_key" binding:"required"`
	Product     string `json:"product" binding:"required"`
	MachineID   string `json:"machine_id" binding:"required"`
	MachineName string `json:"machine_name"`
}

type ValidateLicenceResponse struct {
	Valid     bool      `json:"valid"`
	Reason    string    `json:"reason,omitempty"`
	Tier      string    `json:"tier,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Features  []string  `json:"features,omitempty"`
}

// featuresByTier defines which features are unlocked per product tier.
// Extend this map as products and tiers are added.
var featuresByTier = map[string]map[string][]string{
	"offloadr": {
		"trial": {"single_destination"},
		"pro":   {"unlimited_destinations", "cloud_upload", "provr_manifests", "lto_archive"},
		"studio": {"unlimited_destinations", "cloud_upload", "provr_manifests", "lto_archive", "multi_seat", "priority_support"},
	},
	"scrollr": {
		"trial": {"single_export"},
		"pro":   {"unlimited_exports", "custom_templates", "8k_export"},
		"studio": {"unlimited_exports", "custom_templates", "8k_export", "multi_seat"},
	},
}

// ValidateLicence handles POST /v1/licence/validate
// Called by desktop apps on every launch to verify their licence.
func ValidateLicence(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ValidateLicenceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Background()

		// Look up the licence
		var (
			licenceID    string
			product      string
			tier         string
			status       string
			machineLimit int
			expiresAt    *time.Time
		)
		err := pool.QueryRow(ctx, `
			SELECT id, product, tier, status, machine_limit, expires_at
			FROM licences
			WHERE licence_key = $1
		`, req.LicenceKey).Scan(&licenceID, &product, &tier, &status, &machineLimit, &expiresAt)

		if err != nil {
			c.JSON(http.StatusOK, ValidateLicenceResponse{Valid: false, Reason: "licence_not_found"})
			return
		}

		// Check product matches
		if product != req.Product {
			c.JSON(http.StatusOK, ValidateLicenceResponse{Valid: false, Reason: "product_mismatch"})
			return
		}

		// Check status
		if status != "active" {
			c.JSON(http.StatusOK, ValidateLicenceResponse{Valid: false, Reason: "licence_" + status})
			return
		}

		// Check expiry
		if expiresAt != nil && time.Now().After(*expiresAt) {
			c.JSON(http.StatusOK, ValidateLicenceResponse{Valid: false, Reason: "licence_expired"})
			return
		}

		// Check machine limit — upsert the activation record
		var activationCount int
		_ = pool.QueryRow(ctx, `
			SELECT COUNT(*) FROM licence_activations WHERE licence_id = $1
		`, licenceID).Scan(&activationCount)

		// Try to upsert this machine
		_, upsertErr := pool.Exec(ctx, `
			INSERT INTO licence_activations (licence_id, machine_id, machine_name, last_seen_at)
			VALUES ($1, $2, $3, NOW())
			ON CONFLICT (licence_id, machine_id)
			DO UPDATE SET last_seen_at = NOW(), machine_name = EXCLUDED.machine_name
		`, licenceID, req.MachineID, req.MachineName)

		if upsertErr != nil {
			// Machine is new — check if we're at the limit
			if activationCount >= machineLimit {
				c.JSON(http.StatusOK, ValidateLicenceResponse{Valid: false, Reason: "machine_limit_reached"})
				return
			}
		}

		// Resolve features for this product + tier
		features := []string{}
		if productTiers, ok := featuresByTier[product]; ok {
			if tierFeatures, ok := productTiers[tier]; ok {
				features = tierFeatures
			}
		}

		c.JSON(http.StatusOK, ValidateLicenceResponse{
			Valid:     true,
			Tier:      tier,
			ExpiresAt: expiresAt,
			Features:  features,
		})
	}
}

// ListLicences handles GET /v1/licences (authenticated)
func ListLicences(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.GetString("customer_id")

		rows, err := pool.Query(context.Background(), `
			SELECT id, product, tier, licence_key, status, machine_limit, expires_at, created_at
			FROM licences
			WHERE customer_id = $1
			ORDER BY created_at DESC
		`, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}
		defer rows.Close()

		type LicenceRow struct {
			ID           string     `json:"id"`
			Product      string     `json:"product"`
			Tier         string     `json:"tier"`
			LicenceKey   string     `json:"licence_key"`
			Status       string     `json:"status"`
			MachineLimit int        `json:"machine_limit"`
			ExpiresAt    *time.Time `json:"expires_at"`
			CreatedAt    time.Time  `json:"created_at"`
		}

		var licences []LicenceRow
		for rows.Next() {
			var l LicenceRow
			if err := rows.Scan(&l.ID, &l.Product, &l.Tier, &l.LicenceKey, &l.Status, &l.MachineLimit, &l.ExpiresAt, &l.CreatedAt); err != nil {
				continue
			}
			licences = append(licences, l)
		}

		c.JSON(http.StatusOK, gin.H{"licences": licences})
	}
}

// ListMachines handles GET /v1/licences/:id/machines (authenticated)
func ListMachines(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		licenceID := c.Param("id")
		customerID := c.GetString("customer_id")

		// Verify the licence belongs to this customer
		var count int
		_ = pool.QueryRow(context.Background(), `
			SELECT COUNT(*) FROM licences WHERE id = $1 AND customer_id = $2
		`, licenceID, customerID).Scan(&count)
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "licence not found"})
			return
		}

		rows, err := pool.Query(context.Background(), `
			SELECT id, machine_id, machine_name, activated_at, last_seen_at
			FROM licence_activations
			WHERE licence_id = $1
			ORDER BY last_seen_at DESC
		`, licenceID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}
		defer rows.Close()

		type MachineRow struct {
			ID          string    `json:"id"`
			MachineID   string    `json:"machine_id"`
			MachineName string    `json:"machine_name"`
			ActivatedAt time.Time `json:"activated_at"`
			LastSeenAt  time.Time `json:"last_seen_at"`
		}

		var machines []MachineRow
		for rows.Next() {
			var m MachineRow
			if err := rows.Scan(&m.ID, &m.MachineID, &m.MachineName, &m.ActivatedAt, &m.LastSeenAt); err != nil {
				continue
			}
			machines = append(machines, m)
		}

		c.JSON(http.StatusOK, gin.H{"machines": machines})
	}
}

// DeactivateMachine handles DELETE /v1/licences/:id/machines/:machine_id (authenticated)
func DeactivateMachine(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		licenceID := c.Param("id")
		machineID := c.Param("machine_id")
		customerID := c.GetString("customer_id")

		// Verify ownership
		var count int
		_ = pool.QueryRow(context.Background(), `
			SELECT COUNT(*) FROM licences WHERE id = $1 AND customer_id = $2
		`, licenceID, customerID).Scan(&count)
		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "licence not found"})
			return
		}

		_, err := pool.Exec(context.Background(), `
			DELETE FROM licence_activations WHERE licence_id = $1 AND machine_id = $2
		`, licenceID, machineID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"deactivated": true})
	}
}