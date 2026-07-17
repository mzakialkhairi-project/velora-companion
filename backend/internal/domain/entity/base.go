// Package entity provides the base entity types for the domain layer.
package entity

import "time"

// BaseEntity provides common fields for all domain entities.
// All entities should embed this struct to inherit the base fields.
type BaseEntity struct {
	// ID is the unique identifier using BIGINT auto-increment
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	// CreatedAt is the timestamp when the entity was created
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// UpdatedAt is the timestamp when the entity was last updated
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// DeletedAt is the soft delete timestamp
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName returns the default table name for BaseEntity
func (BaseEntity) TableName() string {
	return ""
}

// Entity is the interface that all domain entities must implement
type Entity interface {
	// GetID returns the entity ID
	GetID() uint64

	// TableName returns the table name for the entity
	TableName() string
}

// GetID implements the Entity interface
func (e BaseEntity) GetID() uint64 {
	return e.ID
}
