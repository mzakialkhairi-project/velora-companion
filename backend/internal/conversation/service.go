// Package conversation provides conversation service functionality.
package conversation

import (
	"context"
	"time"
)

// Service handles conversation business logic.
type Service struct {
	repo Repository
}

// NewService creates a new conversation Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create creates a new conversation from a CreateConversationRequest.
func (s *Service) Create(ctx context.Context, userID uint64, req *CreateConversationRequest) (*Conversation, error) {
	// Validate request
	if err := ValidateCreateRequest(req); err != nil {
		return nil, err
	}

	// Build conversation entity with defaults
	conv := &Conversation{
		WorkspaceID:      req.WorkspaceID,
		Title:            req.Title,
		Provider:         req.Provider,
		Model:            req.Model,
		SystemPrompt:     req.SystemPrompt,
		StreamingEnabled: true,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	// Apply AI settings with defaults
	if req.Temperature != nil {
		conv.Temperature = *req.Temperature
	} else {
		conv.Temperature = 0.7
	}

	if req.TopP != nil {
		conv.TopP = *req.TopP
	} else {
		conv.TopP = 1.0
	}

	if req.MaxTokens != nil {
		conv.MaxTokens = *req.MaxTokens
	} else {
		conv.MaxTokens = 4096
	}

	if req.StreamingEnabled != nil {
		conv.StreamingEnabled = *req.StreamingEnabled
	}

	// Create in repository
	if err := s.repo.Create(ctx, userID, conv); err != nil {
		return nil, err
	}

	return conv, nil
}

// Update updates an existing conversation.
func (s *Service) Update(ctx context.Context, userID uint64, id uint64, req *UpdateConversationRequest) (*Conversation, error) {
	// Validate request
	if err := ValidateUpdateRequest(req); err != nil {
		return nil, err
	}

	// Find existing conversation
	conv, err := s.repo.FindByID(ctx, userID, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	conv.Title = req.Title
	conv.Provider = req.Provider
	conv.Model = req.Model
	conv.SystemPrompt = req.SystemPrompt
	conv.UpdatedAt = time.Now()

	if req.Temperature != nil {
		conv.Temperature = *req.Temperature
	}
	if req.TopP != nil {
		conv.TopP = *req.TopP
	}
	if req.MaxTokens != nil {
		conv.MaxTokens = *req.MaxTokens
	}
	if req.StreamingEnabled != nil {
		conv.StreamingEnabled = *req.StreamingEnabled
	}

	// Save to repository
	if err := s.repo.Update(ctx, userID, conv); err != nil {
		return nil, err
	}

	return conv, nil
}

// Delete soft deletes a conversation.
func (s *Service) Delete(ctx context.Context, userID uint64, id uint64) error {
	return s.repo.Delete(ctx, userID, id)
}

// GetByID retrieves a conversation by ID.
func (s *Service) GetByID(ctx context.Context, userID uint64, id uint64) (*Conversation, error) {
	return s.repo.FindByID(ctx, userID, id)
}

// ListByWorkspace lists all conversations for a workspace.
func (s *Service) ListByWorkspace(ctx context.Context, userID uint64, workspaceID uint64) ([]*Conversation, error) {
	return s.repo.FindByWorkspaceID(ctx, userID, workspaceID)
}
