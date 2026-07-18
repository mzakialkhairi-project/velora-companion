// Package context provides context providers for building chat prompts.
package context

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// DefaultSystemPrompt is the default system prompt.
const DefaultSystemPrompt = "You are a helpful AI assistant."

// SystemProvider provides the system prompt.
type SystemProvider struct {
	prompt string
}

// NewSystemProvider creates a new SystemProvider.
func NewSystemProvider(prompt string) *SystemProvider {
	if prompt == "" {
		prompt = DefaultSystemPrompt
	}
	return &SystemProvider{prompt: prompt}
}

// Build returns the system prompt message.
func (p *SystemProvider) Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error) {
	return []provider.ChatMessage{
		{Role: provider.RoleSystem, Content: p.prompt},
	}, nil
}
