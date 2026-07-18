// Package conversation provides conversation HTTP handlers.
package conversation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles conversation HTTP requests.
type Handler struct {
	svc Service
}

// NewHandler creates a new conversation Handler.
func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

// Create handles POST /api/conversations
func (h *Handler) Create(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conv, err := h.svc.Create(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ToResponse(conv))
}

// Update handles PUT /api/conversations/:id
func (h *Handler) Update(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conv, err := h.svc.Update(c.Request.Context(), userID, id, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToResponse(conv))
}

// Delete handles DELETE /api/conversations/:id
func (h *Handler) Delete(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), userID, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "conversation deleted"})
}

// GetByID handles GET /api/conversations/:id
func (h *Handler) GetByID(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	conv, err := h.svc.GetByID(c.Request.Context(), userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "conversation not found"})
		return
	}

	c.JSON(http.StatusOK, ToResponse(conv))
}

// ListByWorkspace handles GET /api/workspaces/:workspace_id/conversations
func (h *Handler) ListByWorkspace(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	workspaceID, err := strconv.ParseUint(c.Param("workspace_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workspace_id"})
		return
	}

	convs, err := h.svc.ListByWorkspace(c.Request.Context(), userID, workspaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &ConversationListResponse{
		Conversations: ToResponseList(convs),
		Total:         len(convs),
	})
}
