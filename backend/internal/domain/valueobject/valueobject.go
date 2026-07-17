// Package valueobject provides value object types for the domain layer.
package valueobject

import "fmt"

// ValueObject is the base interface for all value objects.
// Value objects are immutable and compared by their values, not by reference.
type ValueObject interface {
	// Equals checks if this value object equals another
	Equals(other ValueObject) bool
}

// Email represents an email address value object
type Email struct {
	value string
}

// NewEmail creates a new Email value object
func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	// Basic email validation
	if len(value) < 3 || value[0] == '@' {
		return nil, fmt.Errorf("invalid email format")
	}
	return &Email{value: value}, nil
}

// Value returns the email value
func (e *Email) Value() string {
	return e.value
}

// Equals implements ValueObject
func (e *Email) Equals(other ValueObject) bool {
	otherEmail, ok := other.(*Email)
	if !ok {
		return false
	}
	return e.value == otherEmail.value
}

// String implements fmt.Stringer
func (e *Email) String() string {
	return e.value
}

// Password represents a password value object
type Password struct {
	hash string
}

// NewPassword creates a new Password value object with the given hash
func NewPassword(hash string) *Password {
	return &Password{hash: hash}
}

// Hash returns the password hash
func (p *Password) Hash() string {
	return p.hash
}

// Equals implements ValueObject
func (p *Password) Equals(other ValueObject) bool {
	otherPassword, ok := other.(*Password)
	if !ok {
		return false
	}
	return p.hash == otherPassword.hash
}

// String implements fmt.Stringer
func (p *Password) String() string {
	return "***"
}
