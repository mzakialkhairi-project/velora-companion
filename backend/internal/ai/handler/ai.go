// Package handler provides AI HTTP handlers.
package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// AIHandler handles AI-related HTTP requests.
type AIHandler struct{}

// NewAIHandler creates a new AIHandler.
func NewAIHandler() *AIHandler {
	return &AIHandler{}
}

// Health handles GET /api/v1/ai/health.
func (h *AIHandler) Health(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	p, ok := provider.Default()
	if !ok {
		shared.ErrorResponse(c, http.StatusServiceUnavailable, "No AI provider available")
		return
	}

	if err := p.Health(ctx); err != nil {
		shared.ErrorResponse(c, http.StatusServiceUnavailable, "AI provider unhealthy: "+err.Error())
		return
	}

	shared.Success(c, http.StatusOK, gin.H{
		"status": "healthy",
	})
}
