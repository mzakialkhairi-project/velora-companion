// Package conversation provides conversation repository interface.
package conversation

import (
	"context"
)

// Repository defines the interface for conversation repository operations.
type Repository interface {
	// Create inserts a new conversation into the repository.
	Create(ctx context.Context, userID uint64, conv *Conversation) error

	// Update updates an existing conversation in the repository.
	Update(ctx context.Context, userID uint64, conv *Conversation) error

	// UpdateSummary updates the summary field of a conversation.
	UpdateSummary(ctx context.Context, conversationID uint64, summary string) error

	// Delete soft deletes a conversation from the repository.
	Delete(ctx context.Context, userID uint64, id uint64) error

	// FindByID retrieves a conversation by its ID for a specific user.
	FindByID(ctx context.Context, userID uint64, id uint64) (*Conversation, error)

	// FindByWorkspaceID retrieves all conversations for a specific workspace.
	FindByWorkspaceID(ctx context.Context, userID uint64, workspaceID uint64) ([]*Conversation, error)
}
