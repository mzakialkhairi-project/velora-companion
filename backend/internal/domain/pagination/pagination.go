// Package pagination provides pagination types for the domain layer.
package pagination

// Page represents a paginated result
type Page[T any] struct {
	Items      []T   `json:"items"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
}

// NewPage creates a new paginated result
func NewPage[T any](items []T, total int64, page, pageSize int) *Page[T] {
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return &Page[T]{
		Items:      items,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// HasNext returns true if there is a next page
func (p *Page[T]) HasNext() bool {
	return p.Page < p.TotalPages
}

// HasPrevious returns true if there is a previous page
func (p *Page[T]) HasPrevious() bool {
	return p.Page > 1
}

// IsFirst returns true if this is the first page
func (p *Page[T]) IsFirst() bool {
	return p.Page == 1
}

// IsLast returns true if this is the last page
func (p *Page[T]) IsLast() bool {
	return p.Page >= p.TotalPages
}

// Offset returns the database offset
func (p *Page[T]) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the page size (for database queries)
func (p *Page[T]) Limit() int {
	return p.PageSize
}

// Cursor represents a cursor for cursor-based pagination
type Cursor struct {
	Value string
}

// Offset returns the offset derived from the cursor
func (c *Cursor) Offset() int {
	// This is a placeholder - actual implementation depends on cursor encoding
	return 0
}
