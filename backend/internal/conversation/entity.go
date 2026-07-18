// Package conversation provides conversation management functionality.
package conversation

import "time"

// Conversation represents a chat conversation.
type Conversation struct {
	ID               uint64    `json:"id"`
	WorkspaceID      uint64    `json:"workspace_id"`
	Title            string    `json:"title"`
	Provider         string    `json:"provider"`
	Model            string    `json:"model"`
	SystemPrompt     string    `json:"system_prompt"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	MaxTokens        int       `json:"max_tokens"`
	StreamingEnabled bool      `json:"streaming_enabled"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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
