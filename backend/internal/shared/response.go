// Package shared provides shared utilities across the application.
package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error information in response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success sends a successful response
func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    http.StatusText(statusCode),
			Message: message,
		},
	})
}

// HandleError handles errors from services and returns appropriate HTTP responses
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch err {
	case apperrors.ErrValidation:
		ErrorResponse(c, http.StatusBadRequest, "Validation error")
	case apperrors.ErrNotFound:
		ErrorResponse(c, http.StatusNotFound, "Resource not found")
	case apperrors.ErrAlreadyExists:
		ErrorResponse(c, http.StatusConflict, "Resource already exists")
	case apperrors.ErrUnauthorized:
		ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	case apperrors.ErrForbidden:
		ErrorResponse(c, http.StatusForbidden, "Forbidden")
	default:
		ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
	}
}
