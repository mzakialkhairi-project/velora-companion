// Package domain provides AI domain entities.
package domain

import "time"

// Conversation represents a chat conversation within a workspace.
type Conversation struct {
	// ID is the unique identifier for the conversation
	ID uint64 `json:"id"`
	// WorkspaceID is the ID of the workspace this conversation belongs to
	WorkspaceID uint64 `json:"workspace_id"`
	// Title is a human-readable title for the conversation
	Title string `json:"title"`
	// Provider is the name of the AI provider (e.g., "ollama", "openai")
	Provider string `json:"provider"`
	// Model is the specific model identifier (e.g., "llama3", "gpt-4")
	Model string `json:"model"`
	// SystemPrompt contains instructions that set the AI's behavior
	SystemPrompt string `json:"system_prompt,omitempty"`
	// Temperature controls randomness (0.0-2.0, lower = more deterministic)
	Temperature float64 `json:"temperature"`
	// TopP controls diversity via nucleus sampling (0.0-1.0)
	TopP float64 `json:"top_p"`
	// MaxTokens limits the maximum number of tokens in responses
	MaxTokens int `json:"max_tokens"`
	// StreamingEnabled indicates whether streaming responses are enabled
	StreamingEnabled bool `json:"streaming_enabled"`
	// Summary is a condensed version of the conversation history
	Summary string `json:"summary,omitempty"`
	// SummaryUpdatedAt is when the summary was last updated
	SummaryUpdatedAt *time.Time `json:"summary_updated_at,omitempty"`
	// CreatedAt is when the conversation was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the conversation was last modified
	UpdatedAt time.Time `json:"updated_at"`
}

// SetDefaults sets default values for AI settings.
func (c *Conversation) SetDefaults() {
	if c.Temperature == 0 {
		c.Temperature = 0.7
	}
	if c.TopP == 0 {
		c.TopP = 1.0
	}
	if c.MaxTokens == 0 {
		c.MaxTokens = 4096
	}
}
