// Package repository provides authentication repository interfaces.
package repository

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// AuthRepository defines the interface for authentication repository operations
type AuthRepository interface {
	// FindUserByEmail retrieves a user by their email address
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)

	// SaveRefreshToken saves a refresh token for a user
	SaveRefreshToken(ctx context.Context, userID uint64, token string, expiresAt int64) error

	// GetRefreshToken retrieves a refresh token and its metadata
	GetRefreshToken(ctx context.Context, token string) (userID uint64, expiresAt int64, err error)

	// DeleteRefreshToken removes a refresh token
	DeleteRefreshToken(ctx context.Context, token string) error

	// DeleteUserRefreshTokens removes all refresh tokens for a user
	DeleteUserRefreshTokens(ctx context.Context, userID uint64) error
}
