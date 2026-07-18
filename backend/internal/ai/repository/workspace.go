// Package repository provides workspace repository interfaces.
package repository

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
)

// WorkspaceRepository defines the interface for workspace repository operations
type WorkspaceRepository interface {
	// Create inserts a new workspace into the repository
	Create(ctx context.Context, userID uint64, workspace *domain.Workspace) error

	// Update updates an existing workspace in the repository
	Update(ctx context.Context, userID uint64, workspace *domain.Workspace) error

	// Delete soft deletes a workspace from the repository
	Delete(ctx context.Context, userID uint64, id uint64) error

	// FindByID retrieves a workspace by its ID for a specific user
	FindByID(ctx context.Context, userID uint64, id uint64) (*domain.Workspace, error)

	// FindByName retrieves a workspace by its name for a specific user
	FindByName(ctx context.Context, userID uint64, name string) (*domain.Workspace, error)

	// List retrieves all workspaces for a specific user
	List(ctx context.Context, userID uint64) ([]*domain.Workspace, error)

	// ExistsByName checks if a workspace with the given name exists for the user
	ExistsByName(ctx context.Context, userID uint64, name string) (bool, error)
}
