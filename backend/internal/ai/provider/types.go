// Package provider provides AI provider abstraction layer.
package provider

// Role represents the role of a message sender in a chat conversation.
type Role string

const (
	// RoleSystem represents a system message (e.g., instructions)
	RoleSystem Role = "system"
	// RoleUser represents a user message
	RoleUser Role = "user"
	// RoleAssistant represents an assistant message
	RoleAssistant Role = "assistant"
	// RoleTool represents a tool message
	RoleTool Role = "tool"
)

// ChatMessage represents a single message in a chat conversation.
type ChatMessage struct {
	// Role is the sender of the message (system, user, assistant, tool)
	Role Role `json:"role"`
	// Content is the text content of the message
	Content string `json:"content"`
	// Name is an optional name for the message sender (for tool messages)
	Name string `json:"name,omitempty"`
	// ToolCallID is the ID of the tool call this message is responding to
	ToolCallID string `json:"tool_call_id,omitempty"`
}

// ToolCall represents a tool call within a message.
type ToolCall struct {
	// ID is the unique identifier for the tool call
	ID string `json:"id"`
	// Name is the name of the tool being called
	Name string `json:"name"`
	// Arguments is the JSON arguments for the tool call
	Arguments string `json:"arguments"`
}

// Tool represents a tool/function that can be called by the AI.
type Tool struct {
	// Type is the type of tool (e.g., "function")
	Type string `json:"type"`
	// Function contains the function definition
	Function FunctionDefinition `json:"function"`
}

// FunctionDefinition defines a callable function.
type FunctionDefinition struct {
	// Name is the name of the function
	Name string `json:"name"`
	// Description describes what the function does
	Description string `json:"description"`
	// Parameters is a JSON Schema for the function parameters
	Parameters string `json:"parameters"`
}

// Usage represents token usage statistics for an API call.
type Usage struct {
	// PromptTokens is the number of tokens in the prompt
	PromptTokens int `json:"prompt_tokens"`
	// CompletionTokens is the number of tokens in the completion
	CompletionTokens int `json:"completion_tokens"`
	// TotalTokens is the total number of tokens
	TotalTokens int `json:"total_tokens"`
}

// ModelInfo represents information about an available AI model.
type ModelInfo struct {
	// ID is the model identifier
	ID string `json:"id"`
	// Name is a human-readable name for the model
	Name string `json:"name"`
	// Description describes the model capabilities
	Description string `json:"description,omitempty"`
	// ContextLength is the maximum context length in tokens
	ContextLength int `json:"context_length,omitempty"`
	// SupportsStreaming indicates if the model supports streaming responses
	SupportsStreaming bool `json:"supports_streaming"`
	// SupportsVision indicates if the model supports image inputs
	SupportsVision bool `json:"supports_vision"`
	// SupportsFunctionCalling indicates if the model supports function calling
	SupportsFunctionCalling bool `json:"supports_function_calling"`
}

// ChatRequest represents a request to the chat API.
type ChatRequest struct {
	// Model is the model identifier to use
	Model string `json:"model"`
	// Messages is the list of conversation messages
	Messages []ChatMessage `json:"messages"`
	// Temperature controls randomness (0-2, lower = more deterministic)
	Temperature float64 `json:"temperature,omitempty"`
	// TopP controls diversity via nucleus sampling
	TopP float64 `json:"top_p,omitempty"`
	// MaxTokens limits the maximum number of tokens in the response
	MaxTokens int `json:"max_tokens,omitempty"`
	// Stop is a list of strings that stop generation when encountered
	Stop []string `json:"stop,omitempty"`
	// Stream enables streaming responses
	Stream bool `json:"stream,omitempty"`
	// Tools is a list of tools available to the model
	Tools []Tool `json:"tools,omitempty"`
	// ToolChoice controls how tools are selected (auto, none, specific)
	ToolChoice string `json:"tool_choice,omitempty"`
}

// ChatResponse represents a response from the chat API.
type ChatResponse struct {
	// ID is the unique identifier for this response
	ID string `json:"id"`
	// Model is the model that generated the response
	Model string `json:"model"`
	// Content is the generated text content
	Content string `json:"content"`
	// Role is the role of the message sender (usually "assistant")
	Role Role `json:"role"`
	// FinishReason explains why generation stopped
	FinishReason string `json:"finish_reason,omitempty"`
	// Usage contains token usage statistics
	Usage Usage `json:"usage,omitempty"`
	// ToolCalls contains any tool calls made during generation
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

// StreamChunk represents a chunk of a streaming response.
type StreamChunk struct {
	// ID is the unique identifier for this response
	ID string `json:"id"`
	// Model is the model generating the response
	Model string `json:"model"`
	// Delta contains the incremental content update
	Delta string `json:"delta"`
	// FinishReason explains why generation stopped (if finished)
	FinishReason string `json:"finish_reason,omitempty"`
	// Usage contains token usage statistics (usually on final chunk)
	Usage *Usage `json:"usage,omitempty"`
	// ToolCalls contains any tool calls made during generation
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// Error contains error information if an error occurred
	Error string `json:"error,omitempty"`
}

// ChatResult represents the result of a chat operation with observability data.
type ChatResult struct {
	// Content is the generated text content
	Content string `json:"content"`
	// Model is the model that generated the response
	Model string `json:"model"`
	// FinishReason explains why generation stopped
	FinishReason string `json:"finish_reason,omitempty"`
	// PromptTokens is the number of tokens in the prompt
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// CompletionTokens is the number of tokens in the completion
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// TotalTokens is the total number of tokens
	TotalTokens int `json:"total_tokens,omitempty"`
	// Duration is the time taken to generate the response in milliseconds
	Duration int64 `json:"duration_ms,omitempty"`
}
