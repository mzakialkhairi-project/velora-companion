// Package validator provides user validation logic.
package validator

import (
	"regexp"
	"strings"

	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"github.com/mzakiaklhairi/velora/internal/modules/user/dto"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateCreateUserRequest validates the CreateUserRequest
func ValidateCreateUserRequest(req *dto.CreateUserRequest) error {
	// Validate name
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return apperrors.ErrValidation
	}

	// Validate email
	email := strings.TrimSpace(req.Email)
	if email == "" {
		return apperrors.ErrValidation
	}
	if !emailRegex.MatchString(email) {
		return apperrors.ErrValidation
	}

	// Validate password
	password := req.Password
	if len(password) < 8 {
		return apperrors.ErrValidation
	}

	return nil
}

// ValidateUpdateUserRequest validates the UpdateUserRequest
func ValidateUpdateUserRequest(req *dto.UpdateUserRequest) error {
	// Validation logic will be implemented here
	return nil
}

// ValidateUserStatus validates if the status is valid
func ValidateUserStatus(status string) bool {
	switch status {
	case "active", "inactive", "banned":
		return true
	default:
		return false
	}
}
