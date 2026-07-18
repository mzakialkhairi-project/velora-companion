// Package entity provides authentication domain entities.
package entity

import (
	"database/sql"
	"time"
)

// RefreshToken represents a refresh token entity
type RefreshToken struct {
	ID        uint64       `json:"id"`
	UserID    uint64       `json:"user_id"`
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt time.Time    `json:"created_at"`
	RevokedAt sql.NullTime `json:"revoked_at,omitempty"`
}

// IsExpired checks if the token is expired
func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

// IsRevoked checks if the token is revoked
func (rt *RefreshToken) IsRevoked() bool {
	return rt.RevokedAt.Valid
}

// IsValid checks if the token is valid (not expired and not revoked)
func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked()
}
