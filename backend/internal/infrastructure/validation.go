package infrastructure

import (
	"regexp"
)

// EmailRegex is the regex pattern for email validation
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail checks if email is valid
func IsValidEmail(email string) bool {
	return EmailRegex.MatchString(email)
}

// MinLength checks if string length is at least min
func MinLength(s string, min int) bool {
	return len(s) >= min
}

// MaxLength checks if string length is at most max
func MaxLength(s string, max int) bool {
	return len(s) <= max
}

// LengthRange checks if string length is within range
func LengthRange(s string, min, max int) bool {
	return MinLength(s, min) && MaxLength(s, max)
}

// IsEmpty checks if string is empty
func IsEmptyString(s string) bool {
	return len(Trim(s)) == 0
}

// Contains checks if string contains substring
func Contains(s, substr string) bool {
	return contains(s, substr)
}

// contains is case-sensitive substring check
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsAt(s, substr))
}

// containsAt checks if s starts with substr
func containsAt(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// InSlice checks if value is in slice
func InSlice[T comparable](v T, slice []T) bool {
	for _, item := range slice {
		if item == v {
			return true
		}
	}
	return false
}
