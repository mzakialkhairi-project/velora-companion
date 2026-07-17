// Package entity provides user domain entities.
package entity

import (
	"github.com/mzakiaklhairi/velora/internal/domain/entity"
)

// UserStatus represents the status of a user
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
)

// User represents a user entity in the domain layer
type User struct {
	entity.BaseEntity

	// Name is the user's display name
	Name string `gorm:"size:255;not null" json:"name"`

	// Email is the user's unique email address
	Email string `gorm:"size:255;uniqueIndex;not null" json:"email"`

	// PasswordHash is the hashed password
	PasswordHash string `gorm:"size:255;not null" json:"-"`

	// Status is the user's account status
	Status UserStatus `gorm:"size:50;default:active" json:"status"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}

// IsActive returns true if the user is active
func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

// IsBanned returns true if the user is banned
func (u *User) IsBanned() bool {
	return u.Status == UserStatusBanned
}
