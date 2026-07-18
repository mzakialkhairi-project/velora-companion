// Package domain provides AI domain entities.
package domain

// AIProfile represents a saved AI configuration profile.
type AIProfile struct {
	// ID is the unique identifier for the profile
	ID uint64 `json:"id"`
	// Name is the display name of the profile
	Name string `json:"name"`
	// Provider is the name of the AI provider (e.g., "ollama", "openai")
	Provider string `json:"provider"`
	// Model is the specific model identifier (e.g., "llama3", "gpt-4")
	Model string `json:"model"`
	// Temperature controls randomness (0-2, lower = more deterministic)
	Temperature float64 `json:"temperature"`
	// MaxTokens limits the maximum number of tokens in responses
	MaxTokens int `json:"max_tokens"`
	// TopP controls diversity via nucleus sampling
	TopP float64 `json:"top_p"`
	// SystemPrompt contains instructions that set the AI's behavior
	SystemPrompt string `json:"system_prompt,omitempty"`
}
