// Package ollama provides Ollama AI provider implementation.
package ollama

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// StreamChunk represents a chunk in a streaming response.
type StreamChunk struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done       bool   `json:"done"`
	DoneReason string `json:"done_reason,omitempty"`
}

// StreamResult holds the final result of a stream.
type StreamResult struct {
	Content      string
	FinishReason string
}

// Stream performs a streaming chat request to Ollama.
func (c *Client) Stream(ctx context.Context, req ChatRequest) (<-chan StreamResult, <-chan error) {
	resultCh := make(chan StreamResult)
	errCh := make(chan error, 1)

	go func() {
		defer close(resultCh)
		defer close(errCh)

		// Create streaming request
		streamReq := ChatRequest{
			Model:  req.Model,
			Stream: true,
		}
		for _, msg := range req.Messages {
			streamReq.Messages = append(streamReq.Messages, msg)
		}

		data, err := json.Marshal(streamReq)
		if err != nil {
			errCh <- fmt.Errorf("failed to marshal request: %w", err)
			return
		}

		url := c.baseURL + "/api/chat"
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(data)))
		if err != nil {
			errCh <- fmt.Errorf("failed to create request: %w", err)
			return
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(httpReq)
		if err != nil {
			errCh <- fmt.Errorf("failed to send request: %w", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			errCh <- fmt.Errorf("ollama API error: status=%d body=%s", resp.StatusCode, string(body))
			return
		}

		// Parse streaming response
		parser := NewStreamParser(resp.Body)
		for {
			chunk, err := parser.Next()
			if err != nil {
				if err == io.EOF {
					return
				}
				errCh <- fmt.Errorf("failed to parse stream: %w", err)
				return
			}

			if chunk.Done {
				resultCh <- StreamResult{
					Content:      chunk.Message.Content,
					FinishReason: chunk.DoneReason,
				}
				return
			}

			resultCh <- StreamResult{
				Content:      chunk.Message.Content,
				FinishReason: "",
			}
		}
	}()

	return resultCh, errCh
}

// StreamParser parses SSE-like streaming responses from Ollama.
type StreamParser struct {
	reader *bufio.Reader
}

// NewStreamParser creates a new StreamParser.
func NewStreamParser(r io.Reader) *StreamParser {
	return &StreamParser{
		reader: bufio.NewReader(r),
	}
}

// Next reads the next chunk from the stream.
func (p *StreamParser) Next() (*StreamChunk, error) {
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Skip empty lines
	if strings.TrimSpace(line) == "" {
		return p.Next()
	}

	var chunk StreamChunk
	if err := json.Unmarshal([]byte(line), &chunk); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chunk: %w", err)
	}

	return &chunk, nil
}

// TimeoutStream performs a streaming request with a timeout.
func (c *Client) TimeoutStream(ctx context.Context, req ChatRequest, timeout time.Duration) (<-chan StreamResult, <-chan error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	resultCh, errCh := c.Stream(ctx, req)

	// Return a wrapper that cancels context when done
	go func() {
		<-resultCh
		cancel()
	}()

	return resultCh, errCh
}
