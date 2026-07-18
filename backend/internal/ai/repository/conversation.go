// Package repository provides conversation and message repository interfaces.
package repository

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
)

// ConversationRepository defines the interface for conversation repository operations
type ConversationRepository interface {
	// Create inserts a new conversation into the repository
	Create(ctx context.Context, userID uint64, workspaceID uint64, conversation *domain.Conversation) error

	// Update updates an existing conversation in the repository
	Update(ctx context.Context, userID uint64, conversation *domain.Conversation) error

	// Delete soft deletes a conversation from the repository
	Delete(ctx context.Context, userID uint64, id uint64) error

	// FindByID retrieves a conversation by its ID for a specific user
	FindByID(ctx context.Context, userID uint64, id uint64) (*domain.Conversation, error)

	// FindByWorkspaceID retrieves all conversations for a specific workspace
	FindByWorkspaceID(ctx context.Context, userID uint64, workspaceID uint64) ([]*domain.Conversation, error)
}

// MessageRepository defines the interface for message repository operations
type MessageRepository interface {
	// Create inserts a new message into the repository
	Create(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, message *domain.Message) error

	// FindByID retrieves a message by its ID for a specific user
	FindByID(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) (*domain.Message, error)

	// List retrieves all messages for a specific conversation
	List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error)

	// Delete soft deletes a message from the repository
	Delete(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) error
}
