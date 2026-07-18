// Package repository provides PostgreSQL implementation for workspace repository.
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"gorm.io/gorm"
)

// PostgresWorkspaceRepository implements WorkspaceRepository using PostgreSQL and GORM
type PostgresWorkspaceRepository struct {
	db *gorm.DB
}

// NewPostgresWorkspaceRepository creates a new PostgresWorkspaceRepository
func NewPostgresWorkspaceRepository(db *gorm.DB) *PostgresWorkspaceRepository {
	return &PostgresWorkspaceRepository{db: db}
}

// Create inserts a new workspace into the repository
func (r *PostgresWorkspaceRepository) Create(ctx context.Context, userID uint64, workspace *domain.Workspace) error {
	workspace.UserID = userID
	workspace.CreatedAt = time.Now()
	workspace.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Create(workspace)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates an existing workspace in the repository
func (r *PostgresWorkspaceRepository) Update(ctx context.Context, userID uint64, workspace *domain.Workspace) error {
	workspace.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).
		Model(workspace).
		Where("id = ? AND user_id = ?", workspace.ID, userID).
		Updates(map[string]interface{}{
			"name":        workspace.Name,
			"description": workspace.Description,
			"updated_at":  workspace.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// Delete soft deletes a workspace from the repository
func (r *PostgresWorkspaceRepository) Delete(ctx context.Context, userID uint64, id uint64) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.Workspace{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// FindByID retrieves a workspace by its ID for a specific user
func (r *PostgresWorkspaceRepository) FindByID(ctx context.Context, userID uint64, id uint64) (*domain.Workspace, error) {
	var workspace domain.Workspace
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", id, userID).
		First(&workspace)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &workspace, nil
}

// FindByName retrieves a workspace by its name for a specific user
func (r *PostgresWorkspaceRepository) FindByName(ctx context.Context, userID uint64, name string) (*domain.Workspace, error) {
	var workspace domain.Workspace
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND name = ? AND deleted_at IS NULL", userID, name).
		First(&workspace)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &workspace, nil
}

// List retrieves all workspaces for a specific user
func (r *PostgresWorkspaceRepository) List(ctx context.Context, userID uint64) ([]*domain.Workspace, error) {
	var workspaces []*domain.Workspace
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Find(&workspaces)
	if result.Error != nil {
		return nil, result.Error
	}
	return workspaces, nil
}

// ExistsByName checks if a workspace with the given name exists for the user
func (r *PostgresWorkspaceRepository) ExistsByName(ctx context.Context, userID uint64, name string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&domain.Workspace{}).
		Where("user_id = ? AND name = ? AND deleted_at IS NULL", userID, name).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
