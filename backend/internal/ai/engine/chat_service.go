// Package engine provides AI chat orchestration.
package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/prompt"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/ai/summary"
	"github.com/mzakiaklhairi/velora/internal/shared"
	"github.com/mzakiaklhairi/velora/internal/shared/scheduler"
)

// ChatService orchestrates the chat flow.
type ChatService struct {
	conversationLoader   ConversationLoader
	messageHistoryLoader MessageHistoryLoader
	messageWriter        MessageWriter
	providerResolver     ProviderResolver
	promptBuilder        *prompt.Builder
	summaryService       *summary.Service
	scheduler            scheduler.Scheduler
}

// NewChatService creates a new ChatService.
func NewChatService(
	conversationLoader ConversationLoader,
	messageHistoryLoader MessageHistoryLoader,
	messageWriter MessageWriter,
	providerResolver ProviderResolver,
	summaryService *summary.Service,
	sched scheduler.Scheduler,
) *ChatService {
	if sched == nil {
		sched = scheduler.NewImmediateScheduler()
	}
	return &ChatService{
		conversationLoader:   conversationLoader,
		messageHistoryLoader: messageHistoryLoader,
		messageWriter:        messageWriter,
		providerResolver:     providerResolver,
		promptBuilder:        prompt.NewBuilder(),
		summaryService:       summaryService,
		scheduler:            sched,
	}
}

// Chat executes the chat flow:
// 1. Validasi Conversation
// 2. Simpan User Message
// 3. Load History
// 4. Build Prompt using Prompt Builder (with Summary)
// 5. Resolve Provider
// 6. Provider.Chat()
// 7. Simpan Assistant Message
// 8. Schedule Summary Job (if threshold exceeded)
// 9. Return ChatResponse
func (s *ChatService) Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	startTime := time.Now()
	var result *provider.ChatResult

	// 1. Validasi Conversation
	conv, err := s.conversationLoader.FindByID(ctx, req.UserID, req.ConversationID)
	if err != nil {
		s.logChat(ctx, req, "", "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("%w: %v", ErrConversationNotFound, err)
	}

	// Verify conversation belongs to the workspace
	if conv.WorkspaceID != req.WorkspaceID {
		s.logChat(ctx, req, "", "", "failed", time.Since(startTime).Milliseconds(), ErrConversationNotFound)
		return nil, ErrConversationNotFound
	}

	// 2. Simpan User Message
	userMsg := &domain.Message{
		Role:           provider.RoleUser,
		Content:        req.Content,
		ConversationID: req.ConversationID,
	}
	if err := s.messageWriter.Create(ctx, req.UserID, req.WorkspaceID, req.ConversationID, userMsg); err != nil {
		s.logChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// 3. Load History (excluding the message we just saved)
	history, err := s.messageHistoryLoader.List(ctx, req.UserID, req.WorkspaceID, req.ConversationID)
	if err != nil {
		s.logChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
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
		s.logChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, err
	}

	// 6. Provider.Chat()
	chatReq := provider.ChatRequest{
		Model:       conv.Model,
		Messages:    messages,
		Temperature: conv.Temperature,
		MaxTokens:   conv.MaxTokens,
	}

	chatResp, err := aiProvider.Chat(ctx, chatReq)
	if err != nil {
		s.logChat(ctx, req, conv.Model, "", "failed", time.Since(startTime).Milliseconds(), err)
		return nil, fmt.Errorf("%w: %v", ErrProviderChatFailed, err)
	}

	// Calculate duration
	duration := time.Since(startTime).Milliseconds()

	// Build ChatResult
	result = &provider.ChatResult{
		Content:          chatResp.Content,
		Model:            chatResp.Model,
		FinishReason:     chatResp.FinishReason,
		PromptTokens:     chatResp.Usage.PromptTokens,
		CompletionTokens: chatResp.Usage.CompletionTokens,
		TotalTokens:      chatResp.Usage.TotalTokens,
		Duration:         duration,
	}

	// 7. Simpan Assistant Message
	assistantMsg := &domain.Message{
		Role:           provider.RoleAssistant,
		Content:        chatResp.Content,
		ConversationID: req.ConversationID,
	}
	if err := s.messageWriter.Create(ctx, req.UserID, req.WorkspaceID, req.ConversationID, assistantMsg); err != nil {
		s.logChat(ctx, req, conv.Model, "", "failed", duration, err)
		return nil, fmt.Errorf("failed to save assistant message: %w", err)
	}

	// 8. Schedule Summary Job (if threshold exceeded)
	// Reload history to get the latest message count including the new ones
	updatedHistory, err := s.messageHistoryLoader.List(ctx, req.UserID, req.WorkspaceID, req.ConversationID)
	if err != nil {
		shared.Warn("failed to reload history for summary", "error", err)
	} else {
		if s.summaryService != nil && s.summaryService.ShouldSummarize(len(updatedHistory)) {
			// Schedule summary job via scheduler (no longer using goroutine directly)
			summaryJob := summary.NewSummaryJob(s.summaryService, req.UserID, req.WorkspaceID, req.ConversationID)
			if err := s.scheduler.Schedule(ctx, summaryJob); err != nil {
				shared.Warn("failed to schedule summary job", "conversation_id", req.ConversationID, "error", err)
			}
		}
	}

	// Log successful chat
	s.logChat(ctx, req, conv.Model, "ollama", "success", duration, nil)

	// 9. Return ChatResponse
	return &ChatResponse{
		MessageID: assistantMsg.ID,
		Content:   chatResp.Content,
		Result:    result,
	}, nil
}

// logChat logs structured chat observability data.
func (s *ChatService) logChat(_ context.Context, req *ChatRequest, model, providerName string, status string, latencyMs int64, err error) {
	logArgs := []any{
		"conversation_id", req.ConversationID,
		"provider", providerName,
		"model", model,
		"latency_ms", latencyMs,
		"status", status,
	}
	if err != nil {
		logArgs = append(logArgs, "error", err.Error())
	}
	shared.Info("AI chat", logArgs...)
}
