// Package engine provides AI chat orchestration.
package engine

// ChatRequest represents a chat request to the engine.
type ChatRequest struct {
	// UserID is the ID of the user making the request.
	UserID uint64
	// WorkspaceID is the ID of the workspace.
	WorkspaceID uint64
	// ConversationID is the ID of the conversation.
	ConversationID uint64
	// Content is the message content from the user.
	Content string
}
