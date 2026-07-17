// Package shared provides shared utilities across the application.
package shared

// GenerateUUID generates a new UUID string
// Note: UUID library deferred until entity creation
func GenerateUUID() string {
	return ""
}

// StringPtr returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}

// StringValue returns string value or empty string from pointer
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool {
	return &b
}

// BoolValue returns bool value or false from pointer
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
