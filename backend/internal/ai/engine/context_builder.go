// Package engine provides AI chat orchestration.
package engine

import (
	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// ContextBuilder builds chat context from conversation and messages.
type ContextBuilder struct{}

// NewContextBuilder creates a new ContextBuilder.
func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{}
}

// Build converts a conversation and its messages into provider.ChatMessage slice.
// If the conversation has a system prompt, it is prepended as the first message.
func (b *ContextBuilder) Build(conv *domain.Conversation, messages []*domain.Message) []provider.ChatMessage {
	result := make([]provider.ChatMessage, 0, len(messages)+1)

	// Add system prompt if present
	if conv.SystemPrompt != "" {
		result = append(result, provider.ChatMessage{
			Role:    provider.RoleSystem,
			Content: conv.SystemPrompt,
		})
	}

	// Convert domain messages to provider messages
	for _, msg := range messages {
		result = append(result, provider.ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return result
}
