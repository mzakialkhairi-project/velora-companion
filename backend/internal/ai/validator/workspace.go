// Package validator provides workspace validation logic.
package validator

import (
	"strings"

	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
)

const (
	MaxWorkspaceNameLength = 100
)

// ValidateCreateWorkspaceRequest validates the CreateWorkspaceRequest
func ValidateCreateWorkspaceRequest(req *dto.CreateWorkspaceRequest) error {
	// Validate name
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return apperrors.ErrValidation
	}
	if len(name) > MaxWorkspaceNameLength {
		return apperrors.ErrValidation
	}

	return nil
}

// ValidateUpdateWorkspaceRequest validates the UpdateWorkspaceRequest
func ValidateUpdateWorkspaceRequest(req *dto.UpdateWorkspaceRequest) error {
	// Validate name if provided
	if req.Name != "" {
		name := strings.TrimSpace(req.Name)
		if name == "" {
			return apperrors.ErrValidation
		}
		if len(name) > MaxWorkspaceNameLength {
			return apperrors.ErrValidation
		}
	}

	return nil
}
