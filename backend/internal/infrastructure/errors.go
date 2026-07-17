package infrastructure

import (
	"errors"
	"net/http"
)

// Error codes
const (
	ErrCodeInternal           = "INTERNAL_ERROR"
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeConflict           = "CONFLICT"
	ErrCodeUnprocessable      = "UNPROCESSABLE_ENTITY"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
)

// AppError represents an application error
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

// New creates a new AppError
func New(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Wrap creates a new AppError wrapping an existing error
func Wrap(code, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

// Predefined errors
var (
	ErrInternal = New(ErrCodeInternal, "Internal server error")
	ErrNotFound = New(ErrCodeNotFound, "Resource not found")
)

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == ErrCodeNotFound
	}
	return false
}

// GetHTTPStatus returns HTTP status code for error code
func GetHTTPStatus(code string) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeUnprocessable:
		return http.StatusUnprocessableEntity
	case ErrCodeServiceUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
