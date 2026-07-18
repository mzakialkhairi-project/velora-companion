// Package summary provides conversation summary functionality.
package summary

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
)

// ConversationUpdater updates a conversation's summary.
type ConversationUpdater interface {
	// UpdateSummary updates the summary fields of a conversation.
	UpdateSummary(ctx context.Context, conversationID uint64, summary string) error
}

// MessageLoader loads messages for a conversation.
type MessageLoader interface {
	// List loads all messages for a conversation.
	List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error)
}

// SummaryProvider calls the AI provider for summarization.
type SummaryProvider interface {
	// Chat calls the AI provider with a single message.
	Chat(ctx context.Context, systemPrompt string, userMessage string) (string, error)
}

// Config holds summary service configuration.
type Config struct {
	// Threshold is the minimum number of messages before summary is triggered.
	Threshold int
	// MessageLimit is the max messages to include in summary generation.
	MessageLimit int
	// LastMessages is the number of recent messages to keep after summary.
	LastMessages int
}
