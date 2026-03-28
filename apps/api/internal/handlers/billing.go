package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/billingportal/session"
)

// CreateBillingPortalSession handles POST /v1/billing/portal (authenticated)
// Returns a Stripe Customer Portal URL for self-service billing management.
func CreateBillingPortalSession(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.GetString("customer_id")

		// Look up the Stripe customer ID
		var stripeCustomerID string
		err := pool.QueryRow(context.Background(), `
			SELECT stripe_customer_id FROM customers WHERE id = $1
		`, customerID).Scan(&stripeCustomerID)
		if err != nil || stripeCustomerID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
			return
		}

		// Stripe Customer Portal sessions require a real Stripe customer ID
		// (not a pending_ placeholder from magic link sign-up before purchase)
		if len(stripeCustomerID) < 4 || stripeCustomerID[:4] != "cus_" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no billing account found — purchase a product first"})
			return
		}

		returnURL := os.Getenv("APP_BASE_URL")
		if returnURL == "" {
			returnURL = "https://app.hijackr.io"
		}
		returnURL += "/dashboard"

		params := &stripe.BillingPortalSessionParams{
			Customer:  stripe.String(stripeCustomerID),
			ReturnURL: stripe.String(returnURL),
		}

		portalSession, err := session.New(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create billing portal session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"url": portalSession.URL})
	}
}