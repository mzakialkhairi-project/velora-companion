// Package service provides authentication service interfaces.
package service

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/modules/auth/dto"
)

// AuthService defines the interface for authentication service operations
type AuthService interface {
	// Register registers a new user
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error)

	// Login authenticates a user and returns tokens
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)

	// Logout invalidates the user's refresh token
	Logout(ctx context.Context, refreshToken string) error

	// RefreshToken generates new access and refresh tokens
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.AuthResponse, error)

	// ValidateToken validates an access token and returns the user ID
	ValidateToken(ctx context.Context, token string) (uint64, error)
}
