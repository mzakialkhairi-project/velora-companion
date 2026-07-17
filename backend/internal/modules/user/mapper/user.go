// Package mapper provides user mapping functions.
package mapper

import (
	"time"

	"github.com/mzakiaklhairi/velora/internal/modules/user/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// ToResponse converts a User entity to a UserResponse DTO
func ToResponse(user *entity.User) *dto.UserResponse {
	if user == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResponseList converts a list of User entities to UserResponse DTOs
func ToResponseList(users []*entity.User) []*dto.UserResponse {
	if users == nil {
		return []*dto.UserResponse{}
	}

	result := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		result[i] = ToResponse(user)
	}

	return result
}
