// Package mapper provides workspace mapping functions.
package mapper

import (
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/dto"
)

// ToResponse converts a Workspace domain to a WorkspaceResponse DTO
func ToResponse(workspace *domain.Workspace) *dto.WorkspaceResponse {
	if workspace == nil {
		return nil
	}

	return &dto.WorkspaceResponse{
		ID:          workspace.ID,
		UserID:      workspace.UserID,
		Name:        workspace.Name,
		Description: workspace.Description,
		CreatedAt:   workspace.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   workspace.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResponseList converts a list of Workspace domains to WorkspaceResponse DTOs
func ToResponseList(workspaces []*domain.Workspace) []*dto.WorkspaceResponse {
	if workspaces == nil {
		return []*dto.WorkspaceResponse{}
	}

	result := make([]*dto.WorkspaceResponse, len(workspaces))
	for i, workspace := range workspaces {
		result[i] = ToResponse(workspace)
	}

	return result
}
