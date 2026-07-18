// Package engine provides AI chat orchestration.
package engine

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// ConversationLoader loads a conversation by ID.
type ConversationLoader interface {
	// FindByID loads a conversation by user ID and conversation ID.
	// Returns ErrConversationNotFound if not found.
	FindByID(ctx context.Context, userID uint64, conversationID uint64) (*domain.Conversation, error)
}

// MessageHistoryLoader loads message history for a conversation.
type MessageHistoryLoader interface {
	// List loads all messages for a conversation.
	List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error)
}

// MessageWriter saves a message.
type MessageWriter interface {
	// Create saves a new message.
	Create(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, message *domain.Message) error
}

// ProviderResolver resolves a provider from the registry.
type ProviderResolver interface {
	// Resolve returns the default provider.
	Resolve() (provider.Provider, error)
}
