// Package context provides context providers for building chat prompts.
package context

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// Builder orchestrates multiple ContextProviders to build chat context.
type Builder struct {
	providers []ContextProvider
}

// NewBuilder creates a new Builder with the given providers.
func NewBuilder(providers ...ContextProvider) *Builder {
	return &Builder{providers: providers}
}

// Add appends a provider to the builder.
func (b *Builder) Add(provider ContextProvider) {
	b.providers = append(b.providers, provider)
}

// Build runs all providers in order and returns the combined messages.
func (b *Builder) Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error) {
	var messages []provider.ChatMessage

	for _, p := range b.providers {
		part, err := p.Build(ctx, req)
		if err != nil {
			return nil, err
		}
		messages = append(messages, part...)
	}

	return messages, nil
}
