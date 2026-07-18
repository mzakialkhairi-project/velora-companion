// Package dto provides conversation and message data transfer objects.
package dto

// CreateConversationRequest represents the request to create a new conversation
type CreateConversationRequest struct {
	Title        string  `json:"title"`
	Provider     string  `json:"provider"`
	Model        string  `json:"model"`
	SystemPrompt string  `json:"system_prompt"`
	Temperature  float64 `json:"temperature"`
	MaxTokens    int     `json:"max_tokens"`
}

// UpdateConversationRequest represents the request to update a conversation
type UpdateConversationRequest struct {
	Title string `json:"title"`
}

// ConversationResponse represents the conversation response
type ConversationResponse struct {
	ID           uint64  `json:"id"`
	WorkspaceID  uint64  `json:"workspace_id"`
	Title        string  `json:"title,omitempty"`
	Provider     string  `json:"provider,omitempty"`
	Model        string  `json:"model,omitempty"`
	SystemPrompt string  `json:"system_prompt,omitempty"`
	Temperature  float64 `json:"temperature"`
	MaxTokens    int     `json:"max_tokens"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// ConversationListResponse represents a list of conversations
type ConversationListResponse struct {
	Conversations []*ConversationResponse `json:"conversations"`
	Total         int                     `json:"total"`
}

// CreateMessageRequest represents the request to create a new message
type CreateMessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// MessageResponse represents the message response
type MessageResponse struct {
	ID             uint64 `json:"id"`
	ConversationID uint64 `json:"conversation_id"`
	Role           string `json:"role"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
}

// MessageListResponse represents a list of messages
type MessageListResponse struct {
	Messages []*MessageResponse `json:"messages"`
	Total    int                `json:"total"`
}
