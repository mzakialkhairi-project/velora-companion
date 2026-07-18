// Package prompt provides prompt building functionality.
package prompt

import "strings"

// Template represents a prompt template.
type Template struct {
	Content string
}

// NewTemplate creates a new template from a string.
func NewTemplate(content string) *Template {
	return &Template{Content: content}
}

// Render renders the template with the given variables.
func (t *Template) Render(vars map[string]string) string {
	result := t.Content
	for key, value := range vars {
		result = strings.ReplaceAll(result, "{{"+key+"}}", value)
	}
	return result
}
