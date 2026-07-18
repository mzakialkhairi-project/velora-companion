// Package context provides context providers for building chat prompts.
package context

import (
	"context"
	"strings"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// SummaryPrefix is the prefix used for summary context.
const SummaryPrefix = "## Previous Conversation Summary\n\n"

// SummaryProvider provides the conversation summary if available.
type SummaryProvider struct{}

// NewSummaryProvider creates a new SummaryProvider.
func NewSummaryProvider() *SummaryProvider {
	return &SummaryProvider{}
}

// Build returns the summary message if conversation has a summary.
func (p *SummaryProvider) Build(ctx context.Context, req BuildContextRequest) ([]provider.ChatMessage, error) {
	if req.Conversation == nil || req.Conversation.Summary == "" {
		return nil, nil
	}

	content := SummaryPrefix + req.Conversation.Summary

	// Check if summary is not too long (truncate if needed)
	if len(content) > 8000 {
		content = content[:8000] + "\n...(truncated)"
	}

	return []provider.ChatMessage{
		{Role: provider.RoleSystem, Content: content},
	}, nil
}

// FormatSummary formats the summary with a prefix.
func FormatSummary(summary string) string {
	if summary == "" {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(SummaryPrefix)
	sb.WriteString(summary)
	return sb.String()
}
