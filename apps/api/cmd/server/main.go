package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v81"

	"github.com/hijackr-dev/hijackr/apps/api/internal/db"
	"github.com/hijackr-dev/hijackr/apps/api/internal/handlers"
	"github.com/hijackr-dev/hijackr/apps/api/internal/middleware"
)

func main() {
	// Load .env in development
	_ = godotenv.Load()

	// Configure Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		log.Fatal("STRIPE_SECRET_KEY is required")
	}

	// Connect to PostgreSQL
	pool, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Run migrations
	if err := db.Migrate(pool); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Set up Gin
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public routes
	v1 := r.Group("/v1")
	{
		// Licence validation — called by desktop apps on launch
		v1.POST("/licence/validate", handlers.ValidateLicence(pool))

		// Stripe webhooks — called by Stripe on payment events
		v1.POST("/webhooks/stripe", handlers.StripeWebhook(pool))

		// Auth — magic link
		v1.POST("/auth/magic-link", handlers.RequestMagicLink(pool))
		v1.GET("/auth/verify", handlers.VerifyMagicLink(pool))
	}

	// Authenticated routes (require valid session token)
	auth := r.Group("/v1")
	auth.Use(middleware.RequireAuth(pool))
	{
		// Customer account
		auth.GET("/me", handlers.GetMe(pool))

		// Licences
		auth.GET("/licences", handlers.ListLicences(pool))

		// Machine management
		auth.GET("/licences/:id/machines", handlers.ListMachines(pool))
		auth.DELETE("/licences/:id/machines/:machine_id", handlers.DeactivateMachine(pool))

		// Stripe Customer Portal redirect
		auth.POST("/billing/portal", handlers.CreateBillingPortalSession(pool))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("hijackr-api listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}