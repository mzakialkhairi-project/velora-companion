package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Default pagination values
const (
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Page     int
	PageSize int
	Offset   int
	Limit    int
}

// PaginationResponse holds pagination metadata for response
type PaginationResponse struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

// GetPaginationParams extracts pagination parameters from query string
func GetPaginationParams(c *gin.Context) PaginationParams {
	page := DefaultPage
	pageSize := DefaultPageSize

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if ps := c.Query("pageSize"); ps != "" {
		if parsed, err := strconv.Atoi(ps); err == nil && parsed > 0 {
			pageSize = parsed
			if pageSize > MaxPageSize {
				pageSize = MaxPageSize
			}
		}
	}

	offset := (page - 1) * pageSize

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
		Limit:    pageSize,
	}
}

// NewPaginationResponse creates a pagination response
func NewPaginationResponse(page, pageSize int, totalItems int64) PaginationResponse {
	totalPages := int(totalItems) / pageSize
	if int(totalItems)%pageSize > 0 {
		totalPages++
	}

	return PaginationResponse{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool               `json:"success"`
	Data       interface{}        `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination"`
}

// SendPaginatedResponse sends a paginated response
func SendPaginatedResponse(c *gin.Context, data interface{}, pagination PaginationResponse) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	})
}
