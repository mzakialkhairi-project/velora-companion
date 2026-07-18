// Package conversation provides conversation management functionality.
package conversation

// CreateConversationRequest represents the request to create a new conversation.
type CreateConversationRequest struct {
	WorkspaceID      uint64   `json:"workspace_id"`
	Title            string   `json:"title"`
	Provider         string   `json:"provider"`
	Model            string   `json:"model"`
	SystemPrompt     string   `json:"system_prompt"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"top_p,omitempty"`
	MaxTokens        *int     `json:"max_tokens,omitempty"`
	StreamingEnabled *bool    `json:"streaming_enabled,omitempty"`
}

// UpdateConversationRequest represents the request to update a conversation.
type UpdateConversationRequest struct {
	Title            string   `json:"title"`
	Provider         string   `json:"provider"`
	Model            string   `json:"model"`
	SystemPrompt     string   `json:"system_prompt"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"top_p,omitempty"`
	MaxTokens        *int     `json:"max_tokens,omitempty"`
	StreamingEnabled *bool    `json:"streaming_enabled,omitempty"`
}

// ConversationResponse represents the conversation response.
type ConversationResponse struct {
	ID               uint64  `json:"id"`
	WorkspaceID      uint64  `json:"workspace_id"`
	Title            string  `json:"title"`
	Provider         string  `json:"provider"`
	Model            string  `json:"model"`
	SystemPrompt     string  `json:"system_prompt"`
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	MaxTokens        int     `json:"max_tokens"`
	StreamingEnabled bool    `json:"streaming_enabled"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// ConversationListResponse represents a list of conversations.
type ConversationListResponse struct {
	Conversations []*ConversationResponse `json:"conversations"`
	Total         int                     `json:"total"`
}
