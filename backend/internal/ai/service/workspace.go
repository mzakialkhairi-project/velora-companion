// Package service provides workspace service interfaces.
package service

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/repository"
	"github.com/mzakiaklhairi/velora/internal/ai/validator"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
)

// WorkspaceService defines the interface for workspace service operations
type WorkspaceService interface {
	// Create creates a new workspace
	Create(ctx context.Context, userID uint64, req *dto.CreateWorkspaceRequest) (*domain.Workspace, error)

	// Update updates an existing workspace
	Update(ctx context.Context, userID uint64, id uint64, req *dto.UpdateWorkspaceRequest) (*domain.Workspace, error)

	// Delete soft deletes a workspace
	Delete(ctx context.Context, userID uint64, id uint64) error

	// GetByID retrieves a workspace by its ID
	GetByID(ctx context.Context, userID uint64, id uint64) (*domain.Workspace, error)

	// List retrieves all workspaces for a user
	List(ctx context.Context, userID uint64) ([]*domain.Workspace, error)
}

// WorkspaceServiceImpl implements WorkspaceService interface
type WorkspaceServiceImpl struct {
	workspaceRepo repository.WorkspaceRepository
}

// NewWorkspaceServiceImpl creates a new WorkspaceServiceImpl
func NewWorkspaceServiceImpl(workspaceRepo repository.WorkspaceRepository) *WorkspaceServiceImpl {
	return &WorkspaceServiceImpl{
		workspaceRepo: workspaceRepo,
	}
}

// Create creates a new workspace
func (s *WorkspaceServiceImpl) Create(ctx context.Context, userID uint64, req *dto.CreateWorkspaceRequest) (*domain.Workspace, error) {
	// Validate request
	if err := validator.ValidateCreateWorkspaceRequest(req); err != nil {
		return nil, err
	}

	// Check if workspace name already exists for this user
	exists, err := s.workspaceRepo.ExistsByName(ctx, userID, req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperrors.ErrAlreadyExists
	}

	// Create workspace domain
	workspace := &domain.Workspace{
		Name:        req.Name,
		Description: req.Description,
	}

	// Save to repository
	if err := s.workspaceRepo.Create(ctx, userID, workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// Update updates an existing workspace
func (s *WorkspaceServiceImpl) Update(ctx context.Context, userID uint64, id uint64, req *dto.UpdateWorkspaceRequest) (*domain.Workspace, error) {
	// Validate request
	if err := validator.ValidateUpdateWorkspaceRequest(req); err != nil {
		return nil, err
	}

	// Get existing workspace
	workspace, err := s.workspaceRepo.FindByID(ctx, userID, id)
	if err != nil {
		return nil, err
	}

	// Check if new name already exists (if name is being changed)
	if req.Name != "" && req.Name != workspace.Name {
		exists, err := s.workspaceRepo.ExistsByName(ctx, userID, req.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, apperrors.ErrAlreadyExists
		}
		workspace.Name = req.Name
	}

	// Update description if provided
	if req.Description != "" {
		workspace.Description = req.Description
	}

	// Save changes
	if err := s.workspaceRepo.Update(ctx, userID, workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// Delete soft deletes a workspace
func (s *WorkspaceServiceImpl) Delete(ctx context.Context, userID uint64, id uint64) error {
	return s.workspaceRepo.Delete(ctx, userID, id)
}

// GetByID retrieves a workspace by its ID
func (s *WorkspaceServiceImpl) GetByID(ctx context.Context, userID uint64, id uint64) (*domain.Workspace, error) {
	return s.workspaceRepo.FindByID(ctx, userID, id)
}

// List retrieves all workspaces for a user
func (s *WorkspaceServiceImpl) List(ctx context.Context, userID uint64) ([]*domain.Workspace, error) {
	return s.workspaceRepo.List(ctx, userID)
}
