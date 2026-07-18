// Package engine provides AI chat orchestration.
package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/prompt"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/shared"
)

// StreamChatService orchestrates the streaming chat flow.
type StreamChatService struct {
	conversationLoader   ConversationLoader
	messageHistoryLoader MessageHistoryLoader
	messageWriter        MessageWriter
	providerResolver     ProviderResolver
	promptBuilder        *prompt.Builder
}

// NewStreamChatService creates a new StreamChatService.
func NewStreamChatService(
	conversationLoader ConversationLoader,
	messageHistoryLoader MessageHistoryLoader,
	messageWriter MessageWriter,
	providerResolver ProviderResolver,
) *StreamChatService {
	return &StreamChatService{
		conversationLoader:   conversationLoader,
		messageHistoryLoader: messageHistoryLoader,
		messageWriter:        messageWriter,
		providerResolver:     providerResolver,
		promptBuilder:        prompt.NewBuilder(),
	}
}

// StreamChatResult contains the result of a streaming chat operation.
type StreamChatResult struct {
	// MessageID is the ID of the saved assistant message.
	MessageID uint64
	// Content is the complete response content from the AI.
	Content string
	// Result contains observability data for the chat operation.
	Result *provider.ChatResult
}

// StreamChat executes the streaming chat flow:
// 1. Validasi Conversation
// 2. Simpan User Message
// 3. Load History
// 4. Build Prompt using Prompt Builder
// 5. Resolve Provider
// 6. Provider.Stream()
// 7. Buffer chunks into complete response
// 8. Simpan Assistant Message (after stream completes)
// 9. Return StreamChatResult
func (s *StreamChatService) StreamChat(ctx context.Context, req *ChatRequest) (*StreamChatResult, error) {
	startTime := time.Now()
	var result *provider.ChatResult

	// 1. Validasi Conversation
	conv, err := s.conversationLoader.FindByID(ctx, req.UserID, req.ConversationID)
	if err != nil {
		s.logStreamChat(ctx, req, "", "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("%w: %v", ErrConversationNotFound, err)
	}

	// Verify conversation belongs to the workspace
	if conv.WorkspaceID != req.WorkspaceID {
		s.logStreamChat(ctx, req, "", "", "failed", time.Since(startTime).Milliseconds(), ErrConversationNotFound)
		return nil, ErrConversationNotFound
	}

	// 2. Simpan User Message
	userMsg := &domain.Message{
		Role:           provider.RoleUser,
		Content:        req.Content,
		ConversationID: req.ConversationID,
	}
	if err := s.messageWriter.Create(ctx, req.UserID, req.WorkspaceID, req.ConversationID, userMsg); err != nil {
		s.logStreamChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// 3. Load History (excluding the message we just saved)
	history, err := s.messageHistoryLoader.List(ctx, req.UserID, req.WorkspaceID, req.ConversationID)
	if err != nil {
		s.logStreamChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("failed to load message history: %w", err)
	}

	// 4. Build Prompt using Prompt Builder
	promptReq := &prompt.BuildRequest{
		Conversation: conv,
		Messages:     history,
		UserMessage:  req.Content,
	}
	messages := s.promptBuilder.Build(promptReq)

	// 5. Resolve Provider
	aiProvider, err := s.providerResolver.Resolve()
	if err != nil {
		s.logStreamChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, err
	}

	// 6. Provider.Stream()
	chatReq := provider.ChatRequest{
		Model:       conv.Model,
		Messages:    messages,
		Temperature: conv.Temperature,
		MaxTokens:   conv.MaxTokens,
	}

	streamCh, err := aiProvider.Stream(ctx, chatReq)
	if err != nil {
		s.logStreamChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("%w: %v", ErrProviderChatFailed, err)
	}

	// 7. Buffer chunks into complete response
	var fullContent strings.Builder
	var finishReason string
	var modelName string

	for chunk := range streamCh {
		if chunk.Error != "" {
			s.logStreamChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), fmt.Errorf("stream error: %s", chunk.Error))
			return nil, fmt.Errorf("%w: %s", ErrProviderChatFailed, chunk.Error)
		}

		fullContent.WriteString(chunk.Delta)
		modelName = chunk.Model
		if chunk.FinishReason != "" {
			finishReason = chunk.FinishReason
		}
	}

	// Calculate duration
	duration := time.Since(startTime).Milliseconds()

	// Build ChatResult
	result = &provider.ChatResult{
		Content:      fullContent.String(),
		Model:        modelName,
		FinishReason: finishReason,
		Duration:     duration,
	}

	// 8. Simpan Assistant Message (after stream completes)
	assistantMsg := &domain.Message{
		Role:           provider.RoleAssistant,
		Content:        fullContent.String(),
		ConversationID: req.ConversationID,
	}
	if err := s.messageWriter.Create(ctx, req.UserID, req.WorkspaceID, req.ConversationID, assistantMsg); err != nil {
		s.logStreamChat(ctx, req, conv.Model, "", "failed", duration, err)
		return nil, fmt.Errorf("failed to save assistant message: %w", err)
	}

	// Log successful stream chat
	s.logStreamChat(ctx, req, conv.Model, "ollama", "success", duration, nil)

	// 9. Return StreamChatResult
	return &StreamChatResult{
		MessageID: assistantMsg.ID,
		Content:   fullContent.String(),
		Result:    result,
	}, nil
}

// logStreamChat logs structured stream chat observability data.
func (s *StreamChatService) logStreamChat(_ context.Context, req *ChatRequest, model, providerName string, status string, latencyMs int64, err error) {
	logArgs := []any{
		"conversation_id", req.ConversationID,
		"provider", providerName,
		"model", model,
		"latency_ms", latencyMs,
		"status", status,
		"stream", true,
	}
	if err != nil {
		logArgs = append(logArgs, "error", err.Error())
	}
	shared.Info("AI stream chat", logArgs...)
}
