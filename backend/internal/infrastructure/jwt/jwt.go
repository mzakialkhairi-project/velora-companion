// Package jwt provides JWT token generation and validation utilities.
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
)

// Claims represents JWT claims
type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Name  string `json:"name"`
}

// JWTService handles JWT operations
type JWTService struct {
	secret        []byte
	issuer        string
	accessExpires time.Duration
}

// Config holds JWT configuration
type Config struct {
	Secret        string
	Issuer        string
	AccessExpires string
}

// NewJWTService creates a new JWT service
func NewJWTService(cfg Config) (*JWTService, error) {
	if cfg.Secret == "" {
		return nil, errors.New("JWT secret is required")
	}

	expires := 24 * time.Hour // default
	if cfg.AccessExpires != "" {
		parsed, err := time.ParseDuration(cfg.AccessExpires)
		if err != nil {
			return nil, err
		}
		expires = parsed
	}

	issuer := cfg.Issuer
	if issuer == "" {
		issuer = "velora"
	}

	return &JWTService{
		secret:        []byte(cfg.Secret),
		issuer:        issuer,
		accessExpires: expires,
	}, nil
}

// GenerateToken generates a JWT token for a user
func (s *JWTService) GenerateToken(userID uint64, email, name string) (string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(s.accessExpires)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", userID),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Email: email,
		Name:  name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		return "", 0, err
	}

	return tokenString, int64(s.accessExpires.Seconds()), nil
}

// ValidateToken validates a JWT token and returns the claims
func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrUnauthorized
		}
		return s.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, apperrors.ErrUnauthorized
		}
		return nil, apperrors.ErrUnauthorized
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, apperrors.ErrUnauthorized
	}

	return claims, nil
}
