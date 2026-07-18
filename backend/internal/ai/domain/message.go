// Package domain provides AI domain entities.
package domain

import (
	"encoding/json"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/ai/types"
)

// Message represents a single message in a conversation.
type Message struct {
	// ID is the unique identifier for the message
	ID uint64 `json:"id"`
	// ConversationID is the ID of the conversation this message belongs to
	ConversationID uint64 `json:"conversation_id"`
	// Role is the sender of the message (system, user, assistant, tool)
	Role provider.Role `json:"role"`
	// Content is the text content of the message
	Content string `json:"content"`
	// Metadata contains additional structured data about the message
	Metadata types.JSONMap `json:"metadata,omitempty"`
	// CreatedAt is when the message was created
	CreatedAt time.Time `json:"created_at"`
}

// GetMetadata retrieves a metadata value by key.
// Returns nil if the key doesn't exist.
func (m *Message) GetMetadata(key string) interface{} {
	if m.Metadata == nil {
		return nil
	}
	return m.Metadata[key]
}

// SetMetadata sets a metadata value by key.
func (m *Message) SetMetadata(key string, value interface{}) {
	if m.Metadata == nil {
		m.Metadata = make(types.JSONMap)
	}
	m.Metadata[key] = value
}

// MetadataToJSON converts metadata to JSON string.
func (m *Message) MetadataToJSON() (string, error) {
	if m.Metadata == nil {
		return "{}", nil
	}
	data, err := json.Marshal(m.Metadata)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromChatMessage creates a Message from a provider.ChatMessage.
func FromChatMessage(msg provider.ChatMessage) *Message {
	return &Message{
		Role:    msg.Role,
		Content: msg.Content,
	}
}

// ToChatMessage converts the Message to a provider.ChatMessage.
func (m *Message) ToChatMessage() provider.ChatMessage {
	return provider.ChatMessage{
		Role:    m.Role,
		Content: m.Content,
	}
}
