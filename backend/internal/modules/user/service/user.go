// Package service provides user service interfaces.
package service

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/modules/user/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// UserService defines the interface for user service operations
type UserService interface {
	// Create creates a new user
	Create(ctx context.Context, req *dto.CreateUserRequest) (*entity.User, error)

	// Update updates an existing user
	Update(ctx context.Context, id uint64, req *dto.UpdateUserRequest) (*entity.User, error)

	// Delete soft deletes a user
	Delete(ctx context.Context, id uint64) error

	// GetByID retrieves a user by their ID
	GetByID(ctx context.Context, id uint64) (*entity.User, error)

	// GetByEmail retrieves a user by their email
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// List retrieves users with pagination
	List(ctx context.Context, page, pageSize int) ([]*entity.User, int64, error)
}
