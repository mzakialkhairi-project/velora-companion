// Package routes provides AI HTTP routes.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/handler"
)

// SetupAIRoutes registers AI routes.
func SetupAIRoutes(r *gin.RouterGroup) {
	aiHandler := handler.NewAIHandler()

	ai := r.Group("/ai")
	{
		ai.GET("/health", aiHandler.Health)
	}
}
