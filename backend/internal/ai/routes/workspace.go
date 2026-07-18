// Package routes provides workspace route definitions.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/handler"
	"github.com/mzakiaklhairi/velora/internal/infrastructure"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
)

// RegisterRoutes registers the workspace routes
func RegisterRoutes(router *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler, jwtService *jwt.JWTService) {
	workspaces := router.Group("/workspaces")
	workspaces.Use(infrastructure.AuthMiddleware(jwtService))
	{
		workspaces.POST("", workspaceHandler.Create)
		workspaces.GET("", workspaceHandler.List)
		workspaces.GET("/:workspaceId", workspaceHandler.GetByID)
		workspaces.PUT("/:workspaceId", workspaceHandler.Update)
		workspaces.DELETE("/:workspaceId", workspaceHandler.Delete)
	}
}
