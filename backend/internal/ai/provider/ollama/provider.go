// Package ollama provides Ollama AI provider implementation.
package ollama

import (
	"context"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// OllamaProvider implements provider.Provider for Ollama.
type OllamaProvider struct {
	client  *Client
	model   string
	timeout time.Duration
}

// NewOllamaProvider creates a new Ollama provider.
func NewOllamaProvider(cfg *Config) *OllamaProvider {
	return &OllamaProvider{
		client:  NewClient(cfg),
		model:   cfg.Model,
		timeout: cfg.Timeout,
	}
}

// Chat implements provider.Provider.Chat.
func (p *OllamaProvider) Chat(ctx context.Context, req provider.ChatRequest) (*provider.ChatResponse, error) {
	// Convert provider.ChatRequest to Ollama request
	ollamaReq := ChatRequest{
		Model:  p.model,
		Stream: false,
	}

	// Map messages
	for _, msg := range req.Messages {
		var role string
		// Map provider.Role to ollama role
		switch msg.Role {
		case provider.RoleSystem:
			role = "system"
		case provider.RoleUser:
			role = "user"
		case provider.RoleAssistant:
			role = "assistant"
		case provider.RoleTool:
			role = "tool"
		default:
			role = "user"
		}

		ollamaReq.Messages = append(ollamaReq.Messages, ChatMessage{
			Role:    role,
			Content: msg.Content,
		})
	}

	// Call Ollama
	ollamaResp, err := p.client.Chat(ctx, ollamaReq)
	if err != nil {
		return nil, err
	}

	// Convert response
	return &provider.ChatResponse{
		ID:           "",
		Model:        ollamaResp.Model,
		Content:      ollamaResp.Message.Content,
		Role:         provider.RoleAssistant,
		FinishReason: "stop",
	}, nil
}

// Stream implements provider.Provider.Stream.
func (p *OllamaProvider) Stream(ctx context.Context, req provider.ChatRequest) (<-chan provider.StreamChunk, error) {
	// Convert provider.ChatRequest to Ollama request
	ollamaReq := ChatRequest{
		Model:  p.model,
		Stream: true,
	}

	// Map messages
	for _, msg := range req.Messages {
		var role string
		switch msg.Role {
		case provider.RoleSystem:
			role = "system"
		case provider.RoleUser:
			role = "user"
		case provider.RoleAssistant:
			role = "assistant"
		case provider.RoleTool:
			role = "tool"
		default:
			role = "user"
		}

		ollamaReq.Messages = append(ollamaReq.Messages, ChatMessage{
			Role:    role,
			Content: msg.Content,
		})
	}

	// Start streaming
	resultCh, errCh := p.client.TimeoutStream(ctx, ollamaReq, p.timeout)

	// Convert to provider.StreamChunk channel
	providerCh := make(chan provider.StreamChunk)

	go func() {
		defer close(providerCh)

		for {
			select {
			case result, ok := <-resultCh:
				if !ok {
					return
				}
				providerCh <- provider.StreamChunk{
					Model:        p.model,
					Delta:        result.Content,
					FinishReason: result.FinishReason,
				}
			case err, ok := <-errCh:
				if ok {
					providerCh <- provider.StreamChunk{
						Error: err.Error(),
					}
				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return providerCh, nil
}

// Models implements provider.Provider.Models.
func (p *OllamaProvider) Models(ctx context.Context) ([]provider.ModelInfo, error) {
	models, err := p.client.ListModels(ctx)
	if err != nil {
		// Return default model if API call fails
		return []provider.ModelInfo{
			{
				ID:                      p.model,
				Name:                    p.model,
				Description:             "Default Ollama model",
				ContextLength:           8192,
				SupportsStreaming:       true,
				SupportsVision:          false,
				SupportsFunctionCalling: false,
			},
		}, nil
	}

	result := make([]provider.ModelInfo, len(models))
	for i, m := range models {
		result[i] = provider.ModelInfo{
			ID:                      m.Name,
			Name:                    m.Name,
			ContextLength:           8192,
			SupportsStreaming:       true,
			SupportsVision:          false,
			SupportsFunctionCalling: false,
		}
	}

	return result, nil
}

// Health implements provider.Provider.Health.
func (p *OllamaProvider) Health(ctx context.Context) error {
	return p.client.Health(ctx)
}
