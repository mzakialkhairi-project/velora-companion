// Package infrastructure provides reusable application infrastructure components.
package infrastructure

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context key types
type contextKey string

const (
	// RequestIDKey is the context key for request ID
	RequestIDKey contextKey = "request_id"
	// RequestIDHeader is the HTTP header for request ID
	RequestIDHeader = "X-Request-ID"
)

// GenerateRequestID generates a random request ID
func GenerateRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// GetRequestID extracts request ID from context
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}

// GetRequestIDFromGin extracts request ID from gin context
func GetRequestIDFromGin(c *gin.Context) string {
	return GetRequestID(c.Request.Context())
}

// RequestIDMiddleware is a middleware that generates or extracts request ID
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get existing request ID from header
		requestID := c.GetHeader(RequestIDHeader)
		if requestID == "" {
			requestID = GenerateRequestID()
		}

		// Store in context
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), RequestIDKey, requestID))

		// Set response header
		c.Header(RequestIDHeader, requestID)

		c.Next()
	}
}

// GetHeader retrieves a header from http.Request
func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}
