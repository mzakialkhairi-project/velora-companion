// Package service provides service interfaces for the domain layer.
package service

import "context"

// Service defines the base service interface
type Service interface {
	// Name returns the service name
	Name() string
}

// CommandHandler defines a handler for a command
type CommandHandler[C, R any] interface {
	// Handle processes a command and returns a result
	Handle(ctx context.Context, cmd C) (R, error)
}

// QueryHandler defines a handler for a query
type QueryHandler[Q, R any] interface {
	// Handle processes a query and returns a result
	Handle(ctx context.Context, query Q) (R, error)
}
