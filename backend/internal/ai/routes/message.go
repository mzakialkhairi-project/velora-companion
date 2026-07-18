// Package routes provides conversation and message route definitions.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/handler"
)

// RegisterConversationRoutes registers conversation routes
func RegisterConversationRoutes(router *gin.RouterGroup, conversationHandler *handler.ConversationHandler) {
	router.POST("/:workspaceId/conversations", conversationHandler.Create)
	router.GET("/:workspaceId/conversations", conversationHandler.List)
	router.GET("/:workspaceId/conversations/:conversationId", conversationHandler.Get)
	router.PUT("/:workspaceId/conversations/:conversationId", conversationHandler.Update)
	router.DELETE("/:workspaceId/conversations/:conversationId", conversationHandler.Delete)
}

// RegisterMessageRoutes registers message routes
func RegisterMessageRoutes(router *gin.RouterGroup, messageHandler *handler.MessageHandler) {
	router.POST("/:workspaceId/conversations/:conversationId/chat", messageHandler.Chat)
	router.POST("/:workspaceId/conversations/:conversationId/messages", messageHandler.Create)
	router.GET("/:workspaceId/conversations/:conversationId/messages", messageHandler.List)
	router.GET("/:workspaceId/conversations/:conversationId/messages/:messageId", messageHandler.Get)
	router.DELETE("/:workspaceId/conversations/:conversationId/messages/:messageId", messageHandler.Delete)
}
