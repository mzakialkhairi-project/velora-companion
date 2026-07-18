// Package ollama provides Ollama AI provider implementation.
package ollama

import (
	"fmt"
	"time"

	"github.com/mzakiaklhairi/velora/internal/shared"
)

const (
	defaultModel   = "qwen2.5:14b"
	defaultTimeout = 120 // seconds
	defaultBaseURL = "http://localhost:11434"
)

// Config holds Ollama configuration.
type Config struct {
	BaseURL string
	Model   string
	Timeout time.Duration
}

// NewConfig creates a new Ollama config from shared config.
func NewConfig(cfg *shared.Config) *Config {
	timeout := defaultTimeout
	if cfg.OllamaTimeout > 0 {
		timeout = cfg.OllamaTimeout
	}

	baseURL := defaultBaseURL
	if cfg.OllamaURL != "" {
		baseURL = cfg.OllamaURL
	}

	model := defaultModel
	if cfg.OllamaModel != "" {
		model = cfg.OllamaModel
	}

	return &Config{
		BaseURL: baseURL,
		Model:   model,
		Timeout: time.Duration(timeout) * time.Second,
	}
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("ollama base URL is required")
	}
	if c.Model == "" {
		return fmt.Errorf("ollama model is required")
	}
	return nil
}
