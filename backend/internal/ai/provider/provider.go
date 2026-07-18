// Package provider provides AI provider abstraction layer.
package provider

import (
	"context"
)

// Provider defines the interface for AI providers.
type Provider interface {
	// Chat sends a chat request and returns a non-streaming response.
	// It blocks until the complete response is received.
	Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)

	// Stream sends a chat request and returns a streaming response channel.
	// The caller should range over the returned channel to receive chunks.
	// The channel is closed when the stream is complete or an error occurs.
	Stream(ctx context.Context, req ChatRequest) (<-chan StreamChunk, error)

	// Models returns a list of available models from this provider.
	Models(ctx context.Context) ([]ModelInfo, error)

	// Health checks if the provider is healthy and accessible.
	// It should verify connectivity and authentication.
	Health(ctx context.Context) error
}
