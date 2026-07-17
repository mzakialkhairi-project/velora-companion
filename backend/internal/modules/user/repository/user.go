// Package repository provides user repository interfaces.
package repository

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	// Create inserts a new user into the repository
	Create(ctx context.Context, user *entity.User) error

	// Update updates an existing user in the repository
	Update(ctx context.Context, user *entity.User) error

	// Delete removes a user from the repository (soft delete)
	Delete(ctx context.Context, id uint64) error

	// FindByID retrieves a user by their ID
	FindByID(ctx context.Context, id uint64) (*entity.User, error)

	// FindByEmail retrieves a user by their email address
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// ExistsByEmail checks if a user with the given email exists
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// List retrieves all users with pagination
	List(ctx context.Context, offset, limit int) ([]*entity.User, error)

	// Count returns the total number of users
	Count(ctx context.Context) (int64, error)
}
