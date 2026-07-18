// Package validator provides conversation and message validation logic.
package validator

import (
	"strings"

	"github.com/mzakiaklhairi/velora/internal/ai/dto"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
)

const (
	MaxMessageContentLength    = 100000
	MaxConversationTitleLength = 100
)

// ValidateCreateConversationRequest validates the CreateConversationRequest
func ValidateCreateConversationRequest(req *dto.CreateConversationRequest) error {
	// Title is optional but if provided, should not be empty
	if req.Title != "" {
		title := strings.TrimSpace(req.Title)
		if title == "" {
			return apperrors.ErrValidation
		}
		if len(title) > MaxConversationTitleLength {
			return apperrors.ErrValidation
		}
	}
	return nil
}

// ValidateUpdateConversationRequest validates the UpdateConversationRequest
func ValidateUpdateConversationRequest(req *dto.UpdateConversationRequest) error {
	if req.Title != "" {
		title := strings.TrimSpace(req.Title)
		if title == "" {
			return apperrors.ErrValidation
		}
		if len(title) > MaxConversationTitleLength {
			return apperrors.ErrValidation
		}
	}
	return nil
}

// ValidateCreateMessageRequest validates the CreateMessageRequest
func ValidateCreateMessageRequest(req *dto.CreateMessageRequest) error {
	// Validate role
	role := strings.TrimSpace(req.Role)
	if role == "" {
		return apperrors.ErrValidation
	}

	// Validate role is one of the allowed values
	switch provider.Role(role) {
	case provider.RoleSystem, provider.RoleUser, provider.RoleAssistant, provider.RoleTool:
		// Valid
	default:
		return apperrors.ErrValidation
	}

	// Validate content
	content := strings.TrimSpace(req.Content)
	if content == "" {
		return apperrors.ErrValidation
	}
	if len(content) > MaxMessageContentLength {
		return apperrors.ErrValidation
	}

	return nil
}

// ValidateRole checks if the role is valid
func ValidateRole(role string) bool {
	switch provider.Role(role) {
	case provider.RoleSystem, provider.RoleUser, provider.RoleAssistant, provider.RoleTool:
		return true
	default:
		return false
	}
}
