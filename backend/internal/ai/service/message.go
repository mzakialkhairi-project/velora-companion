// Package service provides conversation and message service interfaces.
package service

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	"github.com/mzakiaklhairi/velora/internal/ai/repository"
	"github.com/mzakiaklhairi/velora/internal/ai/validator"
)

// ConversationService defines the interface for conversation service operations
type ConversationService interface {
	// Create creates a new conversation
	Create(ctx context.Context, userID uint64, workspaceID uint64, req *dto.CreateConversationRequest) (*domain.Conversation, error)

	// Update updates an existing conversation
	Update(ctx context.Context, userID uint64, id uint64, req *dto.UpdateConversationRequest) (*domain.Conversation, error)

	// Delete soft deletes a conversation
	Delete(ctx context.Context, userID uint64, id uint64) error

	// GetByID retrieves a conversation by its ID
	GetByID(ctx context.Context, userID uint64, id uint64) (*domain.Conversation, error)

	// List retrieves all conversations for a workspace
	List(ctx context.Context, userID uint64, workspaceID uint64) ([]*domain.Conversation, error)
}

// ConversationServiceImpl implements ConversationService interface
type ConversationServiceImpl struct {
	conversationRepo repository.ConversationRepository
}

// NewConversationServiceImpl creates a new ConversationServiceImpl
func NewConversationServiceImpl(conversationRepo repository.ConversationRepository) *ConversationServiceImpl {
	return &ConversationServiceImpl{
		conversationRepo: conversationRepo,
	}
}

// Create creates a new conversation
func (s *ConversationServiceImpl) Create(ctx context.Context, userID uint64, workspaceID uint64, req *dto.CreateConversationRequest) (*domain.Conversation, error) {
	if err := validator.ValidateCreateConversationRequest(req); err != nil {
		return nil, err
	}

	conversation := &domain.Conversation{
		Title:        req.Title,
		Provider:     req.Provider,
		Model:        req.Model,
		SystemPrompt: req.SystemPrompt,
		Temperature:  req.Temperature,
		MaxTokens:    req.MaxTokens,
	}

	if err := s.conversationRepo.Create(ctx, userID, workspaceID, conversation); err != nil {
		return nil, err
	}

	return conversation, nil
}

// Update updates an existing conversation
func (s *ConversationServiceImpl) Update(ctx context.Context, userID uint64, id uint64, req *dto.UpdateConversationRequest) (*domain.Conversation, error) {
	if err := validator.ValidateUpdateConversationRequest(req); err != nil {
		return nil, err
	}

	conversation, err := s.conversationRepo.FindByID(ctx, userID, id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		conversation.Title = req.Title
	}

	if err := s.conversationRepo.Update(ctx, userID, conversation); err != nil {
		return nil, err
	}

	return conversation, nil
}

// Delete soft deletes a conversation
func (s *ConversationServiceImpl) Delete(ctx context.Context, userID uint64, id uint64) error {
	return s.conversationRepo.Delete(ctx, userID, id)
}

// GetByID retrieves a conversation by its ID
func (s *ConversationServiceImpl) GetByID(ctx context.Context, userID uint64, id uint64) (*domain.Conversation, error) {
	return s.conversationRepo.FindByID(ctx, userID, id)
}

// List retrieves all conversations for a workspace
func (s *ConversationServiceImpl) List(ctx context.Context, userID uint64, workspaceID uint64) ([]*domain.Conversation, error) {
	return s.conversationRepo.FindByWorkspaceID(ctx, userID, workspaceID)
}

// MessageService defines the interface for message service operations
type MessageService interface {
	// Create creates a new message
	Create(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, req *dto.CreateMessageRequest) (*domain.Message, error)

	// GetByID retrieves a message by its ID
	GetByID(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) (*domain.Message, error)

	// List retrieves all messages for a conversation
	List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error)

	// Delete soft deletes a message
	Delete(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) error
}

// MessageServiceImpl implements MessageService interface
type MessageServiceImpl struct {
	messageRepo repository.MessageRepository
}

// NewMessageServiceImpl creates a new MessageServiceImpl
func NewMessageServiceImpl(messageRepo repository.MessageRepository) *MessageServiceImpl {
	return &MessageServiceImpl{
		messageRepo: messageRepo,
	}
}

// Create creates a new message
func (s *MessageServiceImpl) Create(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, req *dto.CreateMessageRequest) (*domain.Message, error) {
	if err := validator.ValidateCreateMessageRequest(req); err != nil {
		return nil, err
	}

	message := &domain.Message{
		Role:    provider.Role(req.Role),
		Content: req.Content,
	}

	if err := s.messageRepo.Create(ctx, userID, workspaceID, conversationID, message); err != nil {
		return nil, err
	}

	return message, nil
}

// GetByID retrieves a message by its ID
func (s *MessageServiceImpl) GetByID(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) (*domain.Message, error) {
	return s.messageRepo.FindByID(ctx, userID, workspaceID, conversationID, messageID)
}

// List retrieves all messages for a conversation
func (s *MessageServiceImpl) List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error) {
	return s.messageRepo.List(ctx, userID, workspaceID, conversationID)
}

// Delete soft deletes a message
func (s *MessageServiceImpl) Delete(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) error {
	return s.messageRepo.Delete(ctx, userID, workspaceID, conversationID, messageID)
}
