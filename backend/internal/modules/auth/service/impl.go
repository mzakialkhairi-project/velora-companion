// Package service provides authentication service implementation.
package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"github.com/mzakiaklhairi/velora/internal/infrastructure/jwt"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/entity"
	"github.com/mzakiaklhairi/velora/internal/modules/auth/repository"
	userdto "github.com/mzakiaklhairi/velora/internal/modules/user/dto"
	userentity "github.com/mzakiaklhairi/velora/internal/modules/user/entity"
	userrepo "github.com/mzakiaklhairi/velora/internal/modules/user/repository"
	uservalidator "github.com/mzakiaklhairi/velora/internal/modules/user/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	refreshTokenLength = 64
)

// AuthServiceImpl implements AuthService interface
type AuthServiceImpl struct {
	userRepo         userrepo.UserRepository
	authRepo         repository.AuthRepository
	refreshTokenRepo repository.RefreshTokenRepository
	jwtSvc           *jwt.JWTService
	db               *gorm.DB
	refreshExpires   time.Duration
}

// NewAuthServiceImpl creates a new AuthServiceImpl
func NewAuthServiceImpl(userRepo userrepo.UserRepository, authRepo repository.AuthRepository, refreshTokenRepo repository.RefreshTokenRepository, jwtSvc *jwt.JWTService, db *gorm.DB, refreshExpires time.Duration) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo:         userRepo,
		authRepo:         authRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSvc:           jwtSvc,
		db:               db,
		refreshExpires:   refreshExpires,
	}
}

// generateRefreshToken generates a secure random refresh token
func generateRefreshToken() (string, error) {
	bytes := make([]byte, refreshTokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// Register registers a new user
func (s *AuthServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Validate request
	if err := uservalidator.ValidateCreateUserRequest(&userdto.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		return nil, apperrors.ErrValidation
	}

	// Check if email already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperrors.ErrAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &userentity.User{
		Name:         req.Name,
		Email:        strings.ToLower(req.Email),
		PasswordHash: string(hashedPassword),
		Status:       userentity.UserStatusActive,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Generate access token
	accessToken, expiresIn, err := s.jwtSvc.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Store refresh token
	rt := &entity.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.refreshExpires),
		CreatedAt: time.Now(),
	}
	if err := s.refreshTokenRepo.Create(ctx, rt); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		TokenType:    "Bearer",
	}, nil
}

// Login authenticates a user and returns login response with JWT
func (s *AuthServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Validate request
	if req.Email == "" || req.Password == "" {
		return nil, apperrors.ErrValidation
	}

	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		// Check if it's a not found error
		if err == apperrors.ErrNotFound {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	// Compare password with bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, apperrors.ErrUnauthorized
	}

	// Generate JWT token
	token, expiresIn, err := s.jwtSvc.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Store refresh token
	rt := &entity.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.refreshExpires),
		CreatedAt: time.Now(),
	}
	if err := s.refreshTokenRepo.Create(ctx, rt); err != nil {
		return nil, err
	}

	// Return login response with refresh token
	return &dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
		User: &dto.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Status: string(user.Status),
		},
	}, nil
}

// Logout invalidates the user's refresh token
func (s *AuthServiceImpl) Logout(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		return apperrors.ErrValidation
	}

	// Revoke the refresh token
	if err := s.refreshTokenRepo.RevokeByToken(ctx, refreshToken); err != nil {
		return err
	}

	return nil
}

// RefreshToken generates new access and refresh tokens
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	if req.RefreshToken == "" {
		return nil, apperrors.ErrValidation
	}

	// Find refresh token
	rt, err := s.refreshTokenRepo.FindByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if rt == nil {
		return nil, apperrors.ErrUnauthorized
	}

	// Check if token is valid (not expired and not revoked)
	if !rt.IsValid() {
		return nil, apperrors.ErrUnauthorized
	}

	// Get user information
	user, err := s.userRepo.FindByID(ctx, rt.UserID)
	if err != nil {
		return nil, err
	}

	// Generate new access token
	accessToken, expiresIn, err := s.jwtSvc.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	// Generate new refresh token
	newRefreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Use transaction for rotation
	var newRT *entity.RefreshToken
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Revoke old refresh token
		if err := s.refreshTokenRepo.RevokeByTokenWithTx(ctx, tx, req.RefreshToken); err != nil {
			return err
		}

		// Create new refresh token
		newRT = &entity.RefreshToken{
			UserID:    user.ID,
			Token:     newRefreshToken,
			ExpiresAt: time.Now().Add(s.refreshExpires),
			CreatedAt: time.Now(),
		}
		if err := s.refreshTokenRepo.CreateWithTx(ctx, tx, newRT); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    expiresIn,
		TokenType:    "Bearer",
	}, nil
}

// ValidateToken validates an access token and returns the user ID
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) (uint64, error) {
	claims, err := s.jwtSvc.ValidateToken(token)
	if err != nil {
		return 0, err
	}

	// Parse user ID from subject
	var userID uint64
	if _, err := fmt.Sscanf(claims.Subject, "%d", &userID); err != nil {
		return 0, apperrors.ErrUnauthorized
	}

	return userID, nil
}
