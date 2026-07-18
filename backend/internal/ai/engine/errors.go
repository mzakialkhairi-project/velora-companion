// Package engine provides AI chat orchestration.
package engine

import "errors"

var (
	// ErrConversationNotFound indicates the conversation was not found.
	ErrConversationNotFound = errors.New("conversation not found")

	// ErrProviderNotAvailable indicates the provider is not available.
	ErrProviderNotAvailable = errors.New("provider not available")

	// ErrProviderChatFailed indicates the provider chat call failed.
	ErrProviderChatFailed = errors.New("provider chat failed")

	// ErrInvalidConversation indicates the conversation is invalid.
	ErrInvalidConversation = errors.New("invalid conversation")
)
