// Package context provides context providers for building chat prompts.
package context

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// UserProvider provides the current user message.
type UserProvider struct{}

// NewUserProvider creates a new UserProvider.
func NewUserProvider() *UserProvider {
	return &UserProvider{}
}

// Build returns the current user message.
func (p *UserProvider) Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error) {
	if req.UserMessage == "" {
		return nil, nil
	}

	return []provider.ChatMessage{
		{Role: provider.RoleUser, Content: req.UserMessage},
	}, nil
}
