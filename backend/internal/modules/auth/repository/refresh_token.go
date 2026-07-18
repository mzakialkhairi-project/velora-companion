// Package repository provides authentication repository interfaces and implementations.
package repository

import (
	"context"
	"time"

	"github.com/mzakiaklhairi/velora/internal/modules/auth/entity"
	"gorm.io/gorm"
)

// RefreshTokenRepository defines the interface for refresh token operations
type RefreshTokenRepository interface {
	// Create creates a new refresh token
	Create(ctx context.Context, token *entity.RefreshToken) error

	// FindByToken finds a refresh token by token string
	FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error)

	// RevokeByToken revokes a refresh token by token string
	RevokeByToken(ctx context.Context, token string) error

	// RevokeByUserID revokes all refresh tokens for a user
	RevokeByUserID(ctx context.Context, userID uint64) error

	// DeleteExpired deletes all expired refresh tokens
	DeleteExpired(ctx context.Context) error

	// CreateWithTx creates a new refresh token within a transaction
	CreateWithTx(ctx context.Context, tx *gorm.DB, token *entity.RefreshToken) error

	// RevokeByTokenWithTx revokes a refresh token within a transaction
	RevokeByTokenWithTx(ctx context.Context, tx *gorm.DB, token string) error
}

// PostgresRefreshTokenRepository implements RefreshTokenRepository using PostgreSQL
type PostgresRefreshTokenRepository struct {
	db *gorm.DB
}

// NewPostgresRefreshTokenRepository creates a new PostgresRefreshTokenRepository
func NewPostgresRefreshTokenRepository(db *gorm.DB) *PostgresRefreshTokenRepository {
	return &PostgresRefreshTokenRepository{db: db}
}

// Create creates a new refresh token
func (r *PostgresRefreshTokenRepository) Create(ctx context.Context, token *entity.RefreshToken) error {
	result := r.db.WithContext(ctx).Create(token)
	return result.Error
}

// FindByToken finds a refresh token by token string
func (r *PostgresRefreshTokenRepository) FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	var rt entity.RefreshToken
	result := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(&rt)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rt, nil
}

// RevokeByToken revokes a refresh token by token string
func (r *PostgresRefreshTokenRepository) RevokeByToken(ctx context.Context, token string) error {
	result := r.db.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("token = ? AND revoked_at IS NULL", token).
		Update("revoked_at", time.Now())
	return result.Error
}

// RevokeByUserID revokes all refresh tokens for a user
func (r *PostgresRefreshTokenRepository) RevokeByUserID(ctx context.Context, userID uint64) error {
	result := r.db.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("user_id = ? AND revoked_at IS NULL", userID).
		Update("revoked_at", time.Now())
	return result.Error
}

// DeleteExpired deletes all expired refresh tokens
func (r *PostgresRefreshTokenRepository) DeleteExpired(ctx context.Context) error {
	result := r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now()).
		Delete(&entity.RefreshToken{})
	return result.Error
}

// CreateWithTx creates a new refresh token within a transaction
func (r *PostgresRefreshTokenRepository) CreateWithTx(ctx context.Context, tx *gorm.DB, token *entity.RefreshToken) error {
	result := tx.WithContext(ctx).Create(token)
	return result.Error
}

// RevokeByTokenWithTx revokes a refresh token within a transaction
func (r *PostgresRefreshTokenRepository) RevokeByTokenWithTx(ctx context.Context, tx *gorm.DB, token string) error {
	result := tx.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("token = ? AND revoked_at IS NULL", token).
		Update("revoked_at", time.Now())
	return result.Error
}
