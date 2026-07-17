// Package handler provides user HTTP handlers.
package handler

import (
	"github.com/mzakiaklhairi/velora/internal/modules/user/service"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Create handles user creation
// Create godoc
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) Create() {}

// Update handles user update
// Update godoc
// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRequest true "Update user request"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) Update() {}

// Delete handles user deletion
// Delete godoc
// @Summary Delete a user
// @Description Delete a user account
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) Delete() {}

// GetByID handles getting a user by ID
// GetByID godoc
// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetByID() {}

// List handles listing users with pagination
// List godoc
// @Summary List users
// @Description Get a paginated list of users
// @Tags users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} dto.UserListResponse
// @Router /users [get]
func (h *UserHandler) List() {}
