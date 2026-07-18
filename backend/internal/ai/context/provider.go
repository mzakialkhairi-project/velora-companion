// Package context provides context providers for building chat prompts.
package context

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// BuildContextRequest contains the data needed to build context.
type BuildContextRequest struct {
	Conversation *domain.Conversation
	Messages     []*domain.Message
	UserMessage  string
}

// ContextProvider builds a portion of the chat context.
type ContextProvider interface {
	// Build returns the chat messages for this provider's portion of context.
	Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error)
}
