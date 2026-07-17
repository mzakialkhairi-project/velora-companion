package infrastructure

import (
	"strconv"
	"strings"
)

// ParseID parses a string ID to int64
func ParseID(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// ParseUintID parses a string ID to uint64
func ParseUintID(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// MustParseID parses a string ID to int64, panics on error
func MustParseID(s string) int64 {
	id, err := ParseID(s)
	if err != nil {
		panic(err)
	}
	return id
}

// IsValidID checks if string is a valid positive ID
func IsValidID(s string) bool {
	if s == "" {
		return false
	}
	id, err := ParseID(s)
	return err == nil && id > 0
}

// StringToPtr converts string to pointer
func StringToPtr(s string) *string {
	return &s
}

// StringPtrToValue returns string value or empty string from pointer
func StringPtrToValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// TrimAndLower trims whitespace and converts to lowercase
func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// Trim trims whitespace
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// IsEmpty checks if string is empty or whitespace only
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
