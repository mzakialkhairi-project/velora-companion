// Package mapper provides conversation and message mapping functions.
package mapper

import (
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/dto"
)

// ToConversationResponse converts a Conversation domain to a ConversationResponse DTO
func ToConversationResponse(conversation *domain.Conversation) *dto.ConversationResponse {
	if conversation == nil {
		return nil
	}

	return &dto.ConversationResponse{
		ID:           conversation.ID,
		WorkspaceID:  conversation.WorkspaceID,
		Title:        conversation.Title,
		Provider:     conversation.Provider,
		Model:        conversation.Model,
		SystemPrompt: conversation.SystemPrompt,
		Temperature:  conversation.Temperature,
		MaxTokens:    conversation.MaxTokens,
		CreatedAt:    conversation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    conversation.UpdatedAt.Format(time.RFC3339),
	}
}

// ToConversationResponseList converts a list of Conversation domains to ConversationResponse DTOs
func ToConversationResponseList(conversations []*domain.Conversation) []*dto.ConversationResponse {
	if conversations == nil {
		return []*dto.ConversationResponse{}
	}

	result := make([]*dto.ConversationResponse, len(conversations))
	for i, conversation := range conversations {
		result[i] = ToConversationResponse(conversation)
	}

	return result
}

// ToMessageResponse converts a Message domain to a MessageResponse DTO
func ToMessageResponse(message *domain.Message) *dto.MessageResponse {
	if message == nil {
		return nil
	}

	return &dto.MessageResponse{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		Role:           string(message.Role),
		Content:        message.Content,
		CreatedAt:      message.CreatedAt.Format(time.RFC3339),
	}
}

// ToMessageResponseList converts a list of Message domains to MessageResponse DTOs
func ToMessageResponseList(messages []*domain.Message) []*dto.MessageResponse {
	if messages == nil {
		return []*dto.MessageResponse{}
	}

	result := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		result[i] = ToMessageResponse(message)
	}

	return result
}
