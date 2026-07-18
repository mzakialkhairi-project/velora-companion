// Package summary provides conversation summary functionality.
package summary

import (
	"fmt"
	"strings"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
)

const (
	// DefaultThreshold is the default message count before summary is triggered.
	DefaultThreshold = 100
	// DefaultMessageLimit is the default max messages to include in summary.
	DefaultMessageLimit = 100
	// DefaultLastMessages is the default recent messages to keep after summary.
	DefaultLastMessages = 20
)

// DefaultConfig returns the default summary configuration.
func DefaultConfig() *Config {
	return &Config{
		Threshold:    DefaultThreshold,
		MessageLimit: DefaultMessageLimit,
		LastMessages: DefaultLastMessages,
	}
}

// PromptBuilder builds the summary prompt from messages.
type PromptBuilder struct{}

// NewPromptBuilder creates a new PromptBuilder.
func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{}
}

// BuildMessages builds creates the prompt for summarization.
func (b *PromptBuilder) BuildMessages(messages []*domain.Message, limit int) string {
	var sb strings.Builder
	sb.WriteString("Please summarize the following conversation. ")
	sb.WriteString("Focus on the key topics discussed, important decisions made, ")
	sb.WriteString("and any unresolved issues.\n\n")

	count := 0
	if limit > 0 && len(messages) > limit {
		count = len(messages) - limit
		sb.WriteString("(Previous messages summarized above)\n\n")
	}

	for _, msg := range messages[count:] {
		role := "User"
		if msg.Role == "assistant" {
			role = "Assistant"
		}
		fmt.Fprintf(&sb, "%s: %s\n\n", role, msg.Content)
	}

	return sb.String()
}

// SummarySystemPrompt is the system prompt used for summarization.
const SummarySystemPrompt = `You are a conversation summarizer. Your task is to create a concise summary of the conversation that:

1. Captures the main topics and themes discussed
2. Highlights key information, decisions, or conclusions
3. Notes any important context or background information
4. Identifies any unresolved questions or follow-up items

Keep the summary clear and informative. Focus on what is most important for understanding the conversation's purpose and outcome.`

// BuildSystemPrompt returns the system prompt for summarization.
func (b *PromptBuilder) BuildSystemPrompt() string {
	return SummarySystemPrompt
}
