// Package conversation provides conversation validation functionality.
package conversation

import (
	"fmt"
)

// ValidateCreateRequest validates a CreateConversationRequest.
func ValidateCreateRequest(req *CreateConversationRequest) error {
	if req.WorkspaceID == 0 {
		return fmt.Errorf("workspace_id is required")
	}
	if req.Title == "" {
		return fmt.Errorf("title is required")
	}

	// Validate AI settings
	if err := validateAISettings(req.Temperature, req.TopP, req.MaxTokens); err != nil {
		return err
	}

	return nil
}

// ValidateUpdateRequest validates an UpdateConversationRequest.
func ValidateUpdateRequest(req *UpdateConversationRequest) error {
	if req.Title == "" {
		return fmt.Errorf("title is required")
	}

	// Validate AI settings
	if err := validateAISettings(req.Temperature, req.TopP, req.MaxTokens); err != nil {
		return err
	}

	return nil
}

// validateAISettings validates AI settings for temperature, topP, and maxTokens.
func validateAISettings(temperature, topP *float64, maxTokens *int) error {
	if temperature != nil && (*temperature < 0.0 || *temperature > 2.0) {
		return fmt.Errorf("temperature must be between 0.0 and 2.0")
	}

	if topP != nil && (*topP < 0.0 || *topP > 1.0) {
		return fmt.Errorf("top_p must be between 0.0 and 1.0")
	}

	if maxTokens != nil && *maxTokens < 1 {
		return fmt.Errorf("max_tokens must be at least 1")
	}

	return nil
}
