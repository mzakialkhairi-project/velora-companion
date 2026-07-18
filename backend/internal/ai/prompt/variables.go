// Package prompt provides prompt building functionality.
package prompt

// Variables holds template variables as key-value pairs.
type Variables map[string]string

// NewVariables creates a new Variables instance.
func NewVariables() Variables {
	return make(Variables)
}

// Set sets a variable value.
func (v Variables) Set(key, value string) {
	v[key] = value
}

// Get gets a variable value.
func (v Variables) Get(key string) string {
	return v[key]
}
