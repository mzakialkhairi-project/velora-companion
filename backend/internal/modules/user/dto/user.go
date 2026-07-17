// Package dto provides user data transfer objects.
package dto

import "github.com/mzakiaklhairi/velora/internal/modules/user/entity"

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// UserResponse represents the user response
type UserResponse struct {
	ID        uint64            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Status    entity.UserStatus `json:"status"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// UserListResponse represents a paginated list of users
type UserListResponse struct {
	Users      []*UserResponse `json:"users"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}
