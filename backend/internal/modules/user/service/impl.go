// Package service provides user service implementation.
package service

import (
	"context"

	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"github.com/mzakiaklhairi/velora/internal/modules/user/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
	"github.com/mzakiaklhairi/velora/internal/modules/user/repository"
	"github.com/mzakiaklhairi/velora/internal/modules/user/validator"
)

// UserServiceImpl implements UserService interface
type UserServiceImpl struct {
	userRepo repository.UserRepository
}

// NewUserServiceImpl creates a new UserServiceImpl
func NewUserServiceImpl(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// Create creates a new user
func (s *UserServiceImpl) Create(ctx context.Context, req *dto.CreateUserRequest) (*entity.User, error) {
	// Validate request
	if err := validator.ValidateCreateUserRequest(req); err != nil {
		return nil, err
	}

	// Check if email already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperrors.ErrAlreadyExists
	}

	// Create user entity
	user := &entity.User{
		Name:   req.Name,
		Email:  req.Email,
		Status: entity.UserStatusActive,
	}

	// Note: Password hashing is handled by auth service
	// This service receives the hashed password

	// Save to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Update updates an existing user
func (s *UserServiceImpl) Update(ctx context.Context, id uint64, req *dto.UpdateUserRequest) (*entity.User, error) {
	// Get existing user
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Status != "" {
		user.Status = entity.UserStatus(req.Status)
	}

	// Save changes
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete soft deletes a user
func (s *UserServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.userRepo.Delete(ctx, id)
}

// GetByID retrieves a user by their ID
func (s *UserServiceImpl) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

// GetByEmail retrieves a user by their email
func (s *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

// List retrieves users with pagination
func (s *UserServiceImpl) List(ctx context.Context, page, pageSize int) ([]*entity.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	users, err := s.userRepo.List(ctx, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
