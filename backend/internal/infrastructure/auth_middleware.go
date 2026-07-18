// Package infrastructure provides shared infrastructure components.
package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// AuthMiddleware creates JWT authentication middleware
func AuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		// Check Bearer format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format. Use: Bearer <token>")
			c.Abort()
			return
		}

		tokenString := parts[1]
		if tokenString == "" {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Token is required")
			c.Abort()
			return
		}

		// Validate JWT token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			shared.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Store user info in context
		SetAuthContext(c, claims)

		c.Next()
	}
}
