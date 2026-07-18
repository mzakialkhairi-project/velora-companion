// Package dto provides conversation and message data transfer objects.
package dto

import (
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// ChatRequest represents a chat request body.
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// ChatResponse represents the chat response.
type ChatResponse struct {
	UserMessage      *MessageResponse     `json:"user_message"`
	AssistantMessage *MessageResponse     `json:"assistant_message"`
	Result           *provider.ChatResult `json:"result,omitempty"`
}
