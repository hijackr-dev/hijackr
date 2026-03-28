package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RequireAuth validates the Bearer session token on authenticated routes.
// Sets "customer_id" in the Gin context for downstream handlers.
func RequireAuth(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format — use: Bearer <token>"})
			return
		}

		token := parts[1]

		var customerID string
		var expiresAt time.Time
		err := pool.QueryRow(context.Background(), `
			SELECT customer_id, expires_at FROM sessions WHERE token = $1
		`, token).Scan(&customerID, &expiresAt)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			return
		}

		if time.Now().After(expiresAt) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			return
		}

		c.Set("customer_id", customerID)
		c.Next()
	}
}