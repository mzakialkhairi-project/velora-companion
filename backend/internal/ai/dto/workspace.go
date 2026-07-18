// Package dto provides workspace data transfer objects.
package dto

// CreateWorkspaceRequest represents the request to create a new workspace
type CreateWorkspaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateWorkspaceRequest represents the request to update a workspace
type UpdateWorkspaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WorkspaceResponse represents the workspace response
type WorkspaceResponse struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// WorkspaceListResponse represents a list of workspaces
type WorkspaceListResponse struct {
	Workspaces []*WorkspaceResponse `json:"workspaces"`
	Total      int                  `json:"total"`
}
