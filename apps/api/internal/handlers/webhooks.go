package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

// StripeWebhook handles POST /v1/webhooks/stripe
// Stripe calls this endpoint when payment events occur.
func StripeWebhook(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
			return
		}

		// Verify the webhook signature
		sigHeader := c.GetHeader("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
		event, err := webhook.ConstructEvent(body, sigHeader, webhookSecret)
		if err != nil {
			log.Printf("webhook signature verification failed: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
			return
		}

		switch event.Type {
		case "checkout.session.completed":
			if err := handleCheckoutCompleted(pool, event); err != nil {
				log.Printf("handleCheckoutCompleted error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
				return
			}

		case "customer.subscription.updated":
			if err := handleSubscriptionUpdated(pool, event); err != nil {
				log.Printf("handleSubscriptionUpdated error: %v", err)
			}

		case "customer.subscription.deleted":
			if err := handleSubscriptionDeleted(pool, event); err != nil {
				log.Printf("handleSubscriptionDeleted error: %v", err)
			}

		case "invoice.payment_failed":
			log.Printf("payment failed for event %s", event.ID)
			// TODO: send dunning email
		}

		c.JSON(http.StatusOK, gin.H{"received": true})
	}
}

func handleCheckoutCompleted(pool *pgxpool.Pool, event stripe.Event) error {
	var session stripe.CheckoutSession
	if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
		return fmt.Errorf("unmarshal session: %w", err)
	}

	ctx := context.Background()

	// Upsert customer
	var customerID string
	err := pool.QueryRow(ctx, `
		INSERT INTO customers (stripe_customer_id, email)
		VALUES ($1, $2)
		ON CONFLICT (stripe_customer_id) DO UPDATE SET email = EXCLUDED.email
		RETURNING id
	`, session.Customer.ID, session.CustomerEmail).Scan(&customerID)
	if err != nil {
		return fmt.Errorf("upsert customer: %w", err)
	}

	// Extract product slug from metadata (set in Stripe checkout session)
	product := session.Metadata["product"]
	tier := session.Metadata["tier"]
	if product == "" {
		product = "offloadr" // fallback
	}
	if tier == "" {
		tier = "pro"
	}

	// Generate licence key: XXXX-XXXX-XXXX-XXXX
	licenceKey := generateLicenceKey()

	// Determine expiry (nil = perpetual / managed by subscription status)
	var subscriptionID string
	if session.Subscription != nil {
		subscriptionID = session.Subscription.ID
	}

	_, err = pool.Exec(ctx, `
		INSERT INTO licences (customer_id, product, tier, licence_key, stripe_subscription_id, status)
		VALUES ($1, $2, $3, $4, $5, 'active')
	`, customerID, product, tier, licenceKey, subscriptionID)
	if err != nil {
		return fmt.Errorf("insert licence: %w", err)
	}

	log.Printf("licence created: product=%s tier=%s customer=%s key=%s", product, tier, customerID, licenceKey)
	// TODO: send licence key email to session.CustomerEmail
	return nil
}

func handleSubscriptionUpdated(pool *pgxpool.Pool, event stripe.Event) error {
	var sub stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &sub); err != nil {
		return fmt.Errorf("unmarshal subscription: %w", err)
	}

	status := "active"
	if sub.Status == stripe.SubscriptionStatusCanceled || sub.Status == stripe.SubscriptionStatusUnpaid {
		status = "revoked"
	} else if sub.Status == stripe.SubscriptionStatusPastDue {
		status = "past_due"
	}

	var expiresAt *time.Time
	if sub.CancelAt > 0 {
		t := time.Unix(sub.CancelAt, 0)
		expiresAt = &t
	}

	_, err := pool.Exec(context.Background(), `
		UPDATE licences
		SET status = $1, expires_at = $2, updated_at = NOW()
		WHERE stripe_subscription_id = $3
	`, status, expiresAt, sub.ID)
	return err
}

func handleSubscriptionDeleted(pool *pgxpool.Pool, event stripe.Event) error {
	var sub stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &sub); err != nil {
		return fmt.Errorf("unmarshal subscription: %w", err)
	}

	_, err := pool.Exec(context.Background(), `
		UPDATE licences SET status = 'revoked', updated_at = NOW()
		WHERE stripe_subscription_id = $1
	`, sub.ID)
	return err
}

// generateLicenceKey produces a XXXX-XXXX-XXXX-XXXX format key.
func generateLicenceKey() string {
	id := strings.ToUpper(strings.ReplaceAll(uuid.New().String(), "-", ""))
	return fmt.Sprintf("%s-%s-%s-%s", id[0:4], id[4:8], id[8:12], id[12:16])
}