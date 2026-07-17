// Package validator provides user validation logic.
package validator

import (
	"github.com/mzakiaklhairi/velora/internal/modules/user/dto"
	"github.com/mzakiaklhairi/velora/internal/modules/user/entity"
)

// ValidateCreateUserRequest validates the CreateUserRequest
func ValidateCreateUserRequest(req *dto.CreateUserRequest) error {
	// Validation logic will be implemented here
	return nil
}

// ValidateUpdateUserRequest validates the UpdateUserRequest
func ValidateUpdateUserRequest(req *dto.UpdateUserRequest) error {
	// Validation logic will be implemented here
	return nil
}

// ValidateUserStatus validates if the status is valid
func ValidateUserStatus(status string) bool {
	switch entity.UserStatus(status) {
	case entity.UserStatusActive, entity.UserStatusInactive, entity.UserStatusBanned:
		return true
	default:
		return false
	}
}
