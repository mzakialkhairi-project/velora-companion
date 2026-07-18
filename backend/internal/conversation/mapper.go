// Package conversation provides conversation mapper functionality.
package conversation

import "time"

// ToEntity converts a ConversationResponse to Conversation entity.
func ToEntity(resp *ConversationResponse) *Conversation {
	return &Conversation{
		ID:               resp.ID,
		WorkspaceID:      resp.WorkspaceID,
		Title:            resp.Title,
		Provider:         resp.Provider,
		Model:            resp.Model,
		SystemPrompt:     resp.SystemPrompt,
		Temperature:      resp.Temperature,
		TopP:             resp.TopP,
		MaxTokens:        resp.MaxTokens,
		StreamingEnabled: resp.StreamingEnabled,
		CreatedAt:        parseTime(resp.CreatedAt),
		UpdatedAt:        parseTime(resp.UpdatedAt),
	}
}

// ToResponse converts a Conversation entity to ConversationResponse.
func ToResponse(conv *Conversation) *ConversationResponse {
	return &ConversationResponse{
		ID:               conv.ID,
		WorkspaceID:      conv.WorkspaceID,
		Title:            conv.Title,
		Provider:         conv.Provider,
		Model:            conv.Model,
		SystemPrompt:     conv.SystemPrompt,
		Temperature:      conv.Temperature,
		TopP:             conv.TopP,
		MaxTokens:        conv.MaxTokens,
		StreamingEnabled: conv.StreamingEnabled,
		CreatedAt:        conv.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        conv.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResponseList converts a list of Conversation entities to ConversationResponse list.
func ToResponseList(convs []*Conversation) []*ConversationResponse {
	result := make([]*ConversationResponse, len(convs))
	for i, conv := range convs {
		result[i] = ToResponse(conv)
	}
	return result
}

// parseTime parses a time string in RFC3339 format.
func parseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}
	return t
}
