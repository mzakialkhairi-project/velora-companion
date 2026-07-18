// Package engine provides AI chat orchestration.
package engine

import (
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// ChatResponse represents a chat response from the engine.
type ChatResponse struct {
	// MessageID is the ID of the saved assistant message.
	MessageID uint64
	// Content is the response content from the AI.
	Content string
	// Result contains observability data for the chat operation.
	Result *provider.ChatResult
}
