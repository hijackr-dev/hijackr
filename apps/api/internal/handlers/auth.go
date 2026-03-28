package handlers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RequestMagicLink handles POST /v1/auth/magic-link
// Sends a one-time login link to the customer's email.
func RequestMagicLink(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email string `json:"email" binding:"required,email"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a secure random token
		tokenBytes := make([]byte, 32)
		if _, err := rand.Read(tokenBytes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
			return
		}
		token := hex.EncodeToString(tokenBytes)
		expiresAt := time.Now().Add(15 * time.Minute)

		_, err := pool.Exec(context.Background(), `
			INSERT INTO magic_links (email, token, expires_at)
			VALUES ($1, $2, $3)
		`, req.Email, token, expiresAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// Build the magic link URL
		baseURL := os.Getenv("APP_BASE_URL")
		if baseURL == "" {
			baseURL = "https://app.hijackr.io"
		}
		magicLink := baseURL + "/auth/verify?token=" + token

		// TODO: send email via Resend/Postmark
		// For now, log it in development
		if os.Getenv("ENV") != "production" {
			c.JSON(http.StatusOK, gin.H{
				"message":    "magic link created (dev mode — link returned in response)",
				"magic_link": magicLink,
			})
			return
		}

		// In production, send the email and return a generic response
		_ = magicLink // sendEmail(req.Email, magicLink)
		c.JSON(http.StatusOK, gin.H{"message": "check your email for a login link"})
	}
}

// VerifyMagicLink handles GET /v1/auth/verify?token=...
// Validates the one-time token and issues a session.
func VerifyMagicLink(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
			return
		}

		ctx := context.Background()

		// Look up the magic link
		var (
			linkID    string
			email     string
			expiresAt time.Time
			used      bool
		)
		err := pool.QueryRow(ctx, `
			SELECT id, email, expires_at, used FROM magic_links WHERE token = $1
		`, token).Scan(&linkID, &email, &expiresAt, &used)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		if used || time.Now().After(expiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired or already been used"})
			return
		}

		// Mark token as used
		_, _ = pool.Exec(ctx, `UPDATE magic_links SET used = TRUE WHERE id = $1`, linkID)

		// Upsert customer (create account on first login)
		var customerID string
		err = pool.QueryRow(ctx, `
			INSERT INTO customers (stripe_customer_id, email)
			VALUES ('pending_' || gen_random_uuid()::text, $1)
			ON CONFLICT DO NOTHING
			RETURNING id
		`, email).Scan(&customerID)
		if err != nil || customerID == "" {
			// Customer already exists — look them up
			_ = pool.QueryRow(ctx, `SELECT id FROM customers WHERE email = $1`, email).Scan(&customerID)
		}

		if customerID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not resolve customer"})
			return
		}

		// Create session token
		sessionBytes := make([]byte, 32)
		if _, err := rand.Read(sessionBytes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "session creation failed"})
			return
		}
		sessionToken := hex.EncodeToString(sessionBytes)
		sessionExpiry := time.Now().Add(30 * 24 * time.Hour) // 30 days

		_, err = pool.Exec(ctx, `
			INSERT INTO sessions (customer_id, token, expires_at)
			VALUES ($1, $2, $3)
		`, customerID, sessionToken, sessionExpiry)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "session creation failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"session_token": sessionToken,
			"expires_at":    sessionExpiry,
		})
	}
}

// GetMe handles GET /v1/me (authenticated)
func GetMe(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.GetString("customer_id")

		var email string
		var createdAt time.Time
		err := pool.QueryRow(context.Background(), `
			SELECT email, created_at FROM customers WHERE id = $1
		`, customerID).Scan(&email, &createdAt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":         customerID,
			"email":      email,
			"created_at": createdAt,
		})
	}
}