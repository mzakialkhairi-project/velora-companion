// Package infrastructure provides shared infrastructure components.
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
)

const (
	// authContextKey is the context key for auth context
	authContextKey contextKey = "auth_context"
)

// AuthContext holds authenticated user information
type AuthContext struct {
	UserID uint64
	Email  string
	Name   string
}

// SetAuthContext sets auth context in gin.Context
func SetAuthContext(c *gin.Context, claims *jwt.Claims) {
	// Parse user ID from subject
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)

	authCtx := AuthContext{
		UserID: userID,
		Email:  claims.Email,
		Name:   claims.Name,
	}

	c.Set(string(authContextKey), authCtx)
}

// GetAuthContext gets auth context from gin.Context
func GetAuthContext(c *gin.Context) *AuthContext {
	if val, exists := c.Get(string(authContextKey)); exists {
		if authCtx, ok := val.(AuthContext); ok {
			return &authCtx
		}
	}
	return nil
}

// GetUserID gets user ID from context
func GetUserID(c *gin.Context) uint64 {
	if authCtx := GetAuthContext(c); authCtx != nil {
		return authCtx.UserID
	}
	return 0
}

// GetUserEmail gets user email from context
func GetUserEmail(c *gin.Context) string {
	if authCtx := GetAuthContext(c); authCtx != nil {
		return authCtx.Email
	}
	return ""
}

// GetUserName gets user name from context
func GetUserName(c *gin.Context) string {
	if authCtx := GetAuthContext(c); authCtx != nil {
		return authCtx.Name
	}
	return ""
}

// MustGetUserID gets user ID from context, panics if not found
func MustGetUserID(c *gin.Context) uint64 {
	userID := GetUserID(c)
	if userID == 0 {
		panic("user ID not found in context")
	}
	return userID
}

// MustGetUserEmail gets user email from context, panics if not found
func MustGetUserEmail(c *gin.Context) string {
	email := GetUserEmail(c)
	if email == "" {
		panic("user email not found in context")
	}
	return email
}
