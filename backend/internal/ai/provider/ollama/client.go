// Package ollama provides Ollama AI provider implementation.
package ollama

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the Ollama HTTP client.
type Client struct {
	baseURL    string
	httpClient *http.Client
	timeout    time.Duration
}

// NewClient creates a new Ollama client.
func NewClient(cfg *Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		timeout: cfg.Timeout,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// isNetworkError checks if error is a network error that can be retried.
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	// Check for common network errors
	errStr := err.Error()
	return contains(errStr, "connection refused") ||
		contains(errStr, "connection reset") ||
		contains(errStr, "timeout") ||
		contains(errStr, "temporary failure") ||
		contains(errStr, "i/o timeout")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// doRequest performs an HTTP request to Ollama API with retry logic.
func (c *Client) doRequest(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var lastErr error
	maxRetries := 1 // Retry once on network error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		data, err := c.doRequestOnce(ctx, method, endpoint, body)
		if err == nil {
			return data, nil
		}

		// Only retry on network errors
		if !isNetworkError(err) || attempt == maxRetries {
			return nil, err
		}
		lastErr = err
	}

	return nil, lastErr
}

// doRequestOnce performs a single HTTP request to Ollama API.
func (c *Client) doRequestOnce(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama API error: status=%d body=%s", resp.StatusCode, string(data))
	}

	return data, nil
}

// ChatRequest represents Ollama chat request.
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

// ChatMessage represents a message in Ollama chat request.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatResponse represents Ollama chat response.
type ChatResponse struct {
	Model     string        `json:"model"`
	CreatedAt string        `json:"created_at"`
	Message   OllamaMessage `json:"message"`
	Done      bool          `json:"done"`
}

// OllamaMessage represents the message in Ollama response.
type OllamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Chat performs a chat request to Ollama with context timeout.
func (c *Client) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	// Use context.WithTimeout to enforce timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	data, err := c.doRequest(ctx, http.MethodPost, "/api/chat", req)
	if err != nil {
		return nil, err
	}

	var resp ChatResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// HealthResponse represents Ollama health response.
type HealthResponse struct {
	Status string `json:"status"`
}

// Health checks if Ollama is healthy.
func (c *Client) Health(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	data, err := c.doRequest(ctx, http.MethodGet, "/api/generate", map[string]string{
		"model":  defaultModel,
		"prompt": "test",
	})
	if err != nil {
		return fmt.Errorf("ollama health check failed: %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("ollama returned empty response")
	}

	return nil
}

// TagsResponse represents Ollama tags response for listing models.
type TagsResponse struct {
	Models []ModelInfo `json:"models"`
}

// ModelInfo represents Ollama model info.
type ModelInfo struct {
	Name       string `json:"name"`
	Model      string `json:"model"`
	Size       int64  `json:"size"`
	ModifiedAt string `json:"modified_at"`
}

// ListModels lists available models from Ollama.
func (c *Client) ListModels(ctx context.Context) ([]ModelInfo, error) {
	data, err := c.doRequest(ctx, http.MethodGet, "/api/tags", nil)
	if err != nil {
		return nil, err
	}

	var resp TagsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return resp.Models, nil
}
