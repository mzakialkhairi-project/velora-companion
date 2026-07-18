// Package repository provides PostgreSQL implementation for user repository.
package repository

import (
	"context"
	"errors"
	"time"

	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
	"gorm.io/gorm"
)

// PostgresUserRepository implements UserRepository using PostgreSQL and GORM
type PostgresUserRepository struct {
	db *gorm.DB
}

// NewPostgresUserRepository creates a new PostgresUserRepository
func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// Create inserts a new user into the repository
func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates an existing user in the repository
func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	result := r.db.WithContext(ctx).
		Model(user).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":       user.Name,
			"email":      user.Email,
			"status":     user.Status,
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// Delete soft deletes a user from the repository
func (r *PostgresUserRepository) Delete(ctx context.Context, id uint64) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", id).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// FindByID retrieves a user by their ID
func (r *PostgresUserRepository) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

// FindByEmail retrieves a user by their email address
func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.db.WithContext(ctx).
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *PostgresUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("email = ? AND deleted_at IS NULL", email).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

// List retrieves all users with pagination
func (r *PostgresUserRepository) List(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	var users []*entity.User
	result := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Count returns the total number of users
func (r *PostgresUserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("deleted_at IS NULL").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
