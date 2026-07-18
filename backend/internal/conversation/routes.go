// Package conversation provides conversation routes.
package conversation

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers conversation routes.
func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	conversations := r.Group("/conversations")
	{
		conversations.POST("", h.Create)
		conversations.GET("/:id", h.GetByID)
		conversations.PUT("/:id", h.Update)
		conversations.DELETE("/:id", h.Delete)
	}

	workspaces := r.Group("/workspaces")
	{
		workspaces.GET("/:workspace_id/conversations", h.ListByWorkspace)
	}
}
