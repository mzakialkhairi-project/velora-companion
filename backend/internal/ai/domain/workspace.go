// Package domain provides AI domain entities.
package domain

import "time"

// Workspace represents a user's AI workspace containing conversations.
type Workspace struct {
	// ID is the unique identifier for the workspace
	ID uint64 `json:"id"`
	// UserID is the ID of the user who owns this workspace
	UserID uint64 `json:"user_id"`
	// Name is the display name of the workspace
	Name string `json:"name"`
	// Description provides additional context about the workspace
	Description string `json:"description,omitempty"`
	// CreatedAt is when the workspace was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is when the workspace was last modified
	UpdatedAt time.Time `json:"updated_at"`
}
