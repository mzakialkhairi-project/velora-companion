// Package shared provides shared utilities across the application.
package shared

import (
	"errors"
	"net/http"
)

// Application error codes
const (
	ErrCodeInternal     = "INTERNAL_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeBadRequest   = "BAD_REQUEST"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
)

// AppError represents an application error with code
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError
func NewAppError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common errors
var (
	ErrInternal     = NewAppError(ErrCodeInternal, "Internal server error", nil)
	ErrNotFound     = NewAppError(ErrCodeNotFound, "Resource not found", nil)
	ErrBadRequest   = NewAppError(ErrCodeBadRequest, "Bad request", nil)
	ErrUnauthorized = NewAppError(ErrCodeUnauthorized, "Unauthorized", nil)
	ErrForbidden    = NewAppError(ErrCodeForbidden, "Forbidden", nil)
)

// Database errors
var (
	ErrDBConnection = NewAppError(ErrCodeInternal, "Database connection failed", nil)
	ErrDBQuery      = NewAppError(ErrCodeInternal, "Database query failed", nil)
)

// Redis errors
var (
	ErrRedisConnection = NewAppError(ErrCodeInternal, "Redis connection failed", nil)
	ErrRedisOperation  = NewAppError(ErrCodeInternal, "Redis operation failed", nil)
)

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == ErrCodeNotFound
	}
	return false
}

// GetHTTPStatusCode returns HTTP status code for error
func GetHTTPStatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case ErrCodeNotFound:
			return http.StatusNotFound
		case ErrCodeBadRequest:
			return http.StatusBadRequest
		case ErrCodeUnauthorized:
			return http.StatusUnauthorized
		case ErrCodeForbidden:
			return http.StatusForbidden
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}
