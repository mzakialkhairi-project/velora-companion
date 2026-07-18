// Package prompt provides prompt building functionality.
package prompt

import (
	"context"
	"fmt"
	"strings"

	aicontext "github.com/mzakiaklhairi/velora/internal/ai/context"
	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

const (
	// DefaultLastMessages is the number of recent messages to include.
	DefaultLastMessages = 20
)

// Builder builds chat prompts from conversation and messages.
type Builder struct {
	contextBuilder *aicontext.Builder
}

// NewBuilder creates a new Builder.
func NewBuilder() *Builder {
	// Create context builder with default providers in order:
	// SystemProvider -> SummaryProvider -> HistoryProvider -> UserProvider
	contextBuilder := aicontext.NewBuilder(
		aicontext.NewSystemProvider(DefaultSystemPrompt),
		aicontext.NewSummaryProvider(),
		aicontext.NewHistoryProvider(DefaultLastMessages),
		aicontext.NewUserProvider(),
	)
	return &Builder{contextBuilder: contextBuilder}
}

// BuildRequest contains the input for building a prompt.
type BuildRequest struct {
	Conversation *domain.Conversation
	Messages     []*domain.Message
	UserMessage  string
}

// Build creates a []provider.ChatMessage from the given request.
// The output order is:
// 1. System Prompt (always first)
// 2. Conversation Summary (if available)
// 3. Last N Messages
// 4. Current User Message
func (b *Builder) Build(req *BuildRequest) []provider.ChatMessage {
	ctx := context.Background()

	contextReq := aicontext.BuildContextRequest{
		Conversation: req.Conversation,
		Messages:     req.Messages,
		UserMessage:  req.UserMessage,
	}

	messages, err := b.contextBuilder.Build(ctx, contextReq)
	if err != nil {
		// Fallback: return minimal prompt
		return []provider.ChatMessage{
			{Role: provider.RoleSystem, Content: DefaultSystemPrompt},
			{Role: provider.RoleUser, Content: req.UserMessage},
		}
	}

	return messages
}

// BuildWithMessages creates a prompt with the given messages and user message,
// including summary if available.
func (b *Builder) BuildWithMessages(conversation *domain.Conversation, messages []*domain.Message, userMessage string, lastN int) []provider.ChatMessage {
	req := &BuildRequest{
		Conversation: conversation,
		Messages:     messages,
		UserMessage:  userMessage,
	}

	// Use default Build which already handles lastN via HistoryProvider
	return b.Build(req)
}

// BuildSummaryPrompt builds a prompt for generating a summary.
func (b *Builder) BuildSummaryPrompt(messages []*domain.Message, messageLimit int) (string, string) {
	return SummarySystemPrompt, buildSummaryUserPrompt(messages, messageLimit)
}

// buildSummaryUserPrompt builds the user message for summary generation.
func buildSummaryUserPrompt(messages []*domain.Message, messageLimit int) string {
	var sb strings.Builder
	sb.WriteString("Summarize the following conversation concisely. ")
	sb.WriteString("Capture key topics, important decisions, and any unresolved issues.\n\n")

	count := len(messages)
	if messageLimit > 0 && count > messageLimit {
		count = messageLimit
		sb.WriteString("(Older messages truncated)\n\n")
	}

	for i := 0; i < count; i++ {
		msg := messages[i]
		role := "User"
		if msg.Role == "assistant" {
			role = "Assistant"
		}
		fmt.Fprintf(&sb, "%s: %s\n\n", role, msg.Content)
	}

	return sb.String()
}

// SummarySystemPrompt is used for summary generation.
const SummarySystemPrompt = `You are a conversation summarizer. Create a concise summary that:
1. Captures main topics and themes
2. Highlights key information and decisions
3. Notes important context
4. Identifies unresolved questions

Keep it informative but brief.`
