// Package routes provides authentication route definitions.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/handler"
)

// RegisterRoutes registers the auth routes
func RegisterRoutes(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/refresh", authHandler.RefreshToken)
	}
}
