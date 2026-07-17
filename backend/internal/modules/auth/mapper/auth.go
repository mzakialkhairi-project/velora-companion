// Package mapper provides authentication mapping functions.
package mapper

import (
	"time"

	"github.com/mzakiaklhairi/velora/internal/modules/auth/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// ToAuthResponse converts user entity to AuthResponse DTO
func ToAuthResponse(user *entity.User, accessToken, refreshToken string, expiresIn int64) *dto.AuthResponse {
	if user == nil {
		return nil
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		TokenType:    "Bearer",
	}
}

// ToUserResponse converts user entity to UserResponse DTO
func ToUserResponse(user *entity.User) *dto.UserResponse {
	if user == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    string(user.Status),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
