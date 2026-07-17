package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError represents an error in API response
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success sends a successful response
func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Data:    data,
	})
}

// OK sends HTTP 200 with data
func OK(c *gin.Context, data interface{}) {
	Success(c, http.StatusOK, data)
}

// Created sends HTTP 201 with data
func Created(c *gin.Context, data interface{}) {
	Success(c, http.StatusCreated, data)
}

// NoContent sends HTTP 204
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Fail sends an error response
func Fail(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error: &APIError{
			Code:    code,
			Message: message,
		},
	})
}

// BadRequest sends HTTP 400
func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, "BAD_REQUEST", message)
}

// Unauthorized sends HTTP 401
func Unauthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", message)
}

// Forbidden sends HTTP 403
func Forbidden(c *gin.Context, message string) {
	Fail(c, http.StatusForbidden, "FORBIDDEN", message)
}

// NotFound sends HTTP 404
func NotFound(c *gin.Context, message string) {
	Fail(c, http.StatusNotFound, "NOT_FOUND", message)
}

// Conflict sends HTTP 409
func Conflict(c *gin.Context, message string) {
	Fail(c, http.StatusConflict, "CONFLICT", message)
}

// UnprocessableEntity sends HTTP 422
func UnprocessableEntity(c *gin.Context, message string) {
	Fail(c, http.StatusUnprocessableEntity, "UNPROCESSABLE_ENTITY", message)
}

// InternalError sends HTTP 500
func InternalError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", message)
}

// ServiceUnavailable sends HTTP 503
func ServiceUnavailable(c *gin.Context, message string) {
	Fail(c, http.StatusServiceUnavailable, "SERVICE_UNAVAILABLE", message)
}
