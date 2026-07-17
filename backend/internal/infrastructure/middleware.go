package infrastructure

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware recovers from panics
func RecoveryMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic recovered",
					"error", r,
					"stack", string(debug.Stack()),
					"request_id", GetRequestIDFromGin(c),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, APIResponse{
					Success: false,
					Error: &APIError{
						Code:    ErrCodeInternal,
						Message: "Internal server error",
					},
				})
			}
		}()

		c.Next()
	}
}

// LoggingMiddleware logs requests
func LoggingMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Debug("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"request_id", GetRequestIDFromGin(c),
		)

		c.Next()

		logger.Debug("response",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"request_id", GetRequestIDFromGin(c),
		)
	}
}

// CORSMiddleware handles CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-ID")
		c.Header("Access-Control-Expose-Headers", "X-Request-ID")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
