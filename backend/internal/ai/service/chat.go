// Package service provides conversation and message service interfaces.
package service

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/dto"
)

// ChatService defines the interface for chat operations
type ChatService interface {
	// Chat handles a chat request
	Chat(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, req *dto.ChatRequest) (*dto.ChatResponse, error)
}
