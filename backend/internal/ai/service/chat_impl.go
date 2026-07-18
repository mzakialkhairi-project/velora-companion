// Package service provides conversation and message service interfaces.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/engine"
	"github.com/mzakiaklhairi/velora/internal/ai/mapper"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/ai/repository"
)

// ChatServiceImpl implements ChatService interface using engine.ChatService
type ChatServiceImpl struct {
	conversationRepo repository.ConversationRepository
	messageRepo      repository.MessageRepository
	chatEngine       *engine.ChatService
}

// NewChatServiceImpl creates a new ChatServiceImpl
func NewChatServiceImpl(
	conversationRepo repository.ConversationRepository,
	messageRepo repository.MessageRepository,
	chatEngine *engine.ChatService,
) *ChatServiceImpl {
	return &ChatServiceImpl{
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
		chatEngine:       chatEngine,
	}
}

// Chat handles a chat request
func (s *ChatServiceImpl) Chat(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, req *dto.ChatRequest) (*dto.ChatResponse, error) {
	// Validate request
	if req.Message == "" {
		return nil, fmt.Errorf("message is required")
	}

	// Validate conversation ownership
	conv, err := s.conversationRepo.FindByID(ctx, userID, conversationID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", engine.ErrConversationNotFound, err)
	}
	if conv.WorkspaceID != workspaceID {
		return nil, engine.ErrConversationNotFound
	}

	// Save user message
	userMsg := &domain.Message{
		Role:           provider.RoleUser,
		Content:        req.Message,
		ConversationID: conversationID,
		CreatedAt:      time.Now(),
	}
	if err := s.messageRepo.Create(ctx, userID, workspaceID, conversationID, userMsg); err != nil {
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// Call engine chat service
	engineReq := &engine.ChatRequest{
		UserID:         userID,
		WorkspaceID:    workspaceID,
		ConversationID: conversationID,
		Content:        req.Message,
	}
	engineResp, err := s.chatEngine.Chat(ctx, engineReq)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", engine.ErrProviderChatFailed, err)
	}

	// Get assistant message from repository (saved by engine)
	assistantMsg, err := s.messageRepo.FindByID(ctx, userID, workspaceID, conversationID, engineResp.MessageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get assistant message: %w", err)
	}

	return &dto.ChatResponse{
		UserMessage:      mapper.ToMessageResponse(userMsg),
		AssistantMessage: mapper.ToMessageResponse(assistantMsg),
		Result:           engineResp.Result,
	}, nil
}
