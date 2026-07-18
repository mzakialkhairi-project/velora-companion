// Package context provides context providers for building chat prompts.
package context

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

const (
	// DefaultLastMessages is the default number of recent messages to include.
	DefaultLastMessages = 20
)

// HistoryProvider provides recent conversation history.
type HistoryProvider struct {
	lastN int
}

// NewHistoryProvider creates a new HistoryProvider.
func NewHistoryProvider(lastN int) *HistoryProvider {
	if lastN <= 0 {
		lastN = DefaultLastMessages
	}
	return &HistoryProvider{lastN: lastN}
}

// Build returns the last N messages from history.
func (p *HistoryProvider) Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error) {
	if len(req.Messages) == 0 {
		return nil, nil
	}

	messages := req.Messages

	// If more messages than lastN, take only the last N
	if len(messages) > p.lastN {
		messages = messages[len(messages)-p.lastN:]
	}

	result := make([]provider.ChatMessage, 0, len(messages))
	for _, msg := range messages {
		result = append(result, provider.ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return result, nil
}
