// Package summary provides conversation summary functionality.
package summary

import (
	"context"
	"fmt"

	"github.com/mzakiaklhairi/velora/internal/shared"
)

// Service provides conversation summarization.
type Service struct {
	updater       ConversationUpdater
	loader        MessageLoader
	provider      SummaryProvider
	promptBuilder *PromptBuilder
	config        *Config
}

// NewService creates a new summary Service.
func NewService(
	updater ConversationUpdater,
	loader MessageLoader,
	provider SummaryProvider,
	config *Config,
) *Service {
	if config == nil {
		config = DefaultConfig()
	}
	return &Service{
		updater:       updater,
		loader:        loader,
		provider:      provider,
		promptBuilder: NewPromptBuilder(),
		config:        config,
	}
}

// ShouldSummarize checks if a conversation needs summarization.
func (s *Service) ShouldSummarize(messageCount int) bool {
	return messageCount >= s.config.Threshold
}

// Summarize generates and stores a summary for a conversation.
func (s *Service) Summarize(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) error {
	// Load all messages
	messages, err := s.loader.List(ctx, userID, workspaceID, conversationID)
	if err != nil {
		return fmt.Errorf("failed to load messages: %w", err)
	}

	// Check threshold
	if !s.ShouldSummarize(len(messages)) {
		shared.Debug("skipping summary", "conversation_id", conversationID, "message_count", len(messages))
		return nil
	}

	shared.Info("generating summary", "conversation_id", conversationID, "message_count", len(messages))

	// Build prompt
	systemPrompt := s.promptBuilder.BuildSystemPrompt()
	userPrompt := s.promptBuilder.BuildMessages(messages, s.config.MessageLimit)

	// Call AI provider
	summary, err := s.provider.Chat(ctx, systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf("failed to generate summary: %w", err)
	}

	// Save summary
	if err := s.updater.UpdateSummary(ctx, conversationID, summary); err != nil {
		return fmt.Errorf("failed to save summary: %w", err)
	}

	shared.Info("summary saved", "conversation_id", conversationID, "summary_length", len(summary))
	return nil
}
