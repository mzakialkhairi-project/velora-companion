// Package handler provides conversation and message HTTP handlers.
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/mapper"
	"github.com/mzakiaklhairi/velora/internal/ai/service"
	"github.com/mzakiaklhairi/velora/internal/infrastructure"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// ConversationHandler handles conversation-related HTTP requests
type ConversationHandler struct {
	conversationService service.ConversationService
}

// NewConversationHandler creates a new ConversationHandler
func NewConversationHandler(conversationService service.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		conversationService: conversationService,
	}
}

// Create handles POST /api/v1/workspaces/:workspaceId/conversations
func (h *ConversationHandler) Create(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	var req dto.CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	conversation, err := h.conversationService.Create(c.Request.Context(), userID, workspaceID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusCreated, mapper.ToConversationResponse(conversation))
}

// List handles GET /api/v1/workspaces/:workspaceId/conversations
func (h *ConversationHandler) List(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversations, err := h.conversationService.List(c.Request.Context(), userID, workspaceID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, dto.ConversationListResponse{
		Conversations: mapper.ToConversationResponseList(conversations),
		Total:         len(conversations),
	})
}

// Get handles GET /api/v1/workspaces/:workspaceId/conversations/:id
func (h *ConversationHandler) Get(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	conversation, err := h.conversationService.GetByID(c.Request.Context(), userID, conversationID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, mapper.ToConversationResponse(conversation))
}

// Update handles PUT /api/v1/workspaces/:workspaceId/conversations/:id
func (h *ConversationHandler) Update(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	var req dto.UpdateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	conversation, err := h.conversationService.Update(c.Request.Context(), userID, conversationID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, mapper.ToConversationResponse(conversation))
}

// Delete handles DELETE /api/v1/workspaces/:workspaceId/conversations/:id
func (h *ConversationHandler) Delete(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	if err := h.conversationService.Delete(c.Request.Context(), userID, conversationID); err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusNoContent, nil)
}

// MessageHandler handles message-related HTTP requests
type MessageHandler struct {
	messageService service.MessageService
	chatService    service.ChatService
}

// NewMessageHandler creates a new MessageHandler
func NewMessageHandler(messageService service.MessageService, chatService service.ChatService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
		chatService:    chatService,
	}
}

// Create handles POST /api/v1/workspaces/:workspaceId/conversations/:conversationId/messages
func (h *MessageHandler) Create(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	var req dto.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	message, err := h.messageService.Create(c.Request.Context(), userID, workspaceID, conversationID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusCreated, mapper.ToMessageResponse(message))
}

// List handles GET /api/v1/workspaces/:workspaceId/conversations/:conversationId/messages
func (h *MessageHandler) List(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	messages, err := h.messageService.List(c.Request.Context(), userID, workspaceID, conversationID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, dto.MessageListResponse{
		Messages: mapper.ToMessageResponseList(messages),
		Total:    len(messages),
	})
}

// Get handles GET /api/v1/workspaces/:workspaceId/conversations/:conversationId/messages/:id
func (h *MessageHandler) Get(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	messageID, err := strconv.ParseUint(c.Param("messageId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid message ID")
		return
	}

	message, err := h.messageService.GetByID(c.Request.Context(), userID, workspaceID, conversationID, messageID)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, mapper.ToMessageResponse(message))
}

// Delete handles DELETE /api/v1/workspaces/:workspaceId/conversations/:conversationId/messages/:id
func (h *MessageHandler) Delete(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	messageID, err := strconv.ParseUint(c.Param("messageId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid message ID")
		return
	}

	if err := h.messageService.Delete(c.Request.Context(), userID, workspaceID, conversationID, messageID); err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusNoContent, nil)
}

// Chat handles POST /api/v1/workspaces/:workspaceId/conversations/:conversationId/chat
func (h *MessageHandler) Chat(c *gin.Context) {
	userID := infrastructure.GetUserID(c)

	workspaceID, err := strconv.ParseUint(c.Param("workspaceId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid workspace ID")
		return
	}

	conversationID, err := strconv.ParseUint(c.Param("conversationId"), 10, 64)
	if err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	var req dto.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		shared.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	resp, err := h.chatService.Chat(c.Request.Context(), userID, workspaceID, conversationID, &req)
	if err != nil {
		shared.HandleError(c, err)
		return
	}

	shared.Success(c, http.StatusOK, resp)
}
