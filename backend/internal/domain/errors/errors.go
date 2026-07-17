// Package errors provides domain-level error types that are used across the application.
package errors

import "errors"

// Sentinel errors for common domain scenarios
var (
	// ErrNotFound indicates the requested resource was not found
	ErrNotFound = errors.New("resource not found")

	// ErrAlreadyExists indicates a resource with the same identifier already exists
	ErrAlreadyExists = errors.New("resource already exists")

	// ErrValidation indicates input validation failed
	ErrValidation = errors.New("validation error")

	// ErrConflict indicates a business rule conflict
	ErrConflict = errors.New("conflict")

	// ErrUnauthorized indicates the user is not authenticated
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden indicates the user does not have permission
	ErrForbidden = errors.New("forbidden")
)

// DomainError represents a domain-level error with additional context
type DomainError struct {
	Err     error
	Message string
	Code    string
}

func (e *DomainError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func (e *DomainError) Unwrap() error {
	return e.Err
}

// NewDomainError creates a new domain error
func NewDomainError(err error, message, code string) *DomainError {
	return &DomainError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}
