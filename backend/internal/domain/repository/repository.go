// Package repository provides repository interfaces for the domain layer.
package repository

import (
	"context"

	"github.com/mzakiaklhairi/velora/internal/domain/entity"
	"github.com/mzakiaklhairi/velora/internal/domain/specification"
)

// Repository defines the base repository interface that all repositories should implement.
type Repository[T entity.Entity] interface {
	// Create inserts a new entity into the repository
	Create(ctx context.Context, e T) error

	// Update updates an existing entity in the repository
	Update(ctx context.Context, e T) error

	// Delete removes an entity from the repository (soft delete)
	Delete(ctx context.Context, id uint64) error

	// FindByID retrieves an entity by its ID
	FindByID(ctx context.Context, id uint64) (T, error)

	// Exists checks if an entity with the given ID exists
	Exists(ctx context.Context, id uint64) (bool, error)

	// Count returns the total number of entities
	Count(ctx context.Context) (int64, error)
}

// ReadOnlyRepository defines read-only operations for a repository
type ReadOnlyRepository[T entity.Entity] interface {
	// FindByID retrieves an entity by its ID
	FindByID(ctx context.Context, id uint64) (T, error)

	// Exists checks if an entity with the given ID exists
	Exists(ctx context.Context, id uint64) (bool, error)

	// Count returns the total number of entities
	Count(ctx context.Context) (int64, error)

	// FindAll retrieves all entities
	FindAll(ctx context.Context) ([]T, error)

	// FindBySpec retrieves entities matching the given specification
	FindBySpec(ctx context.Context, spec specification.Specification[T]) ([]T, error)
}

// WriteOnlyRepository defines write-only operations for a repository
type WriteOnlyRepository[T entity.Entity] interface {
	// Create inserts a new entity into the repository
	Create(ctx context.Context, e T) error

	// Update updates an existing entity in the repository
	Update(ctx context.Context, e T) error

	// Delete removes an entity from the repository (soft delete)
	Delete(ctx context.Context, id uint64) error
}
