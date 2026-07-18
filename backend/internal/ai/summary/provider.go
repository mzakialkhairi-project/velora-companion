// Package summary provides conversation summary functionality.
package summary

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// ProviderAdapter wraps an AI provider for summary generation.
type ProviderAdapter struct {
	aiProvider provider.Provider
}

// NewProviderAdapter creates a new ProviderAdapter.
func NewProviderAdapter(aiProvider provider.Provider) *ProviderAdapter {
	return &ProviderAdapter{aiProvider: aiProvider}
}

// Chat calls the AI provider with a summary prompt.
func (p *ProviderAdapter) Chat(ctx context.Context, systemPrompt string, userMessage string) (string, error) {
	messages := []provider.ChatMessage{
		{Role: provider.RoleSystem, Content: systemPrompt},
		{Role: provider.RoleUser, Content: userMessage},
	}

	resp, err := p.aiProvider.Chat(ctx, provider.ChatRequest{
		Messages: messages,
	})
	if err != nil {
		return "", err
	}

	return resp.Content, nil
}
