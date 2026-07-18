// Package repository provides PostgreSQL implementation for conversation and message repositories.
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mzakiaklhairi/velora/internal/ai/domain"
	apperrors "github.com/mzakiaklhairi/velora/internal/domain/errors"
	"gorm.io/gorm"
)

// PostgresConversationRepository implements ConversationRepository using PostgreSQL and GORM
type PostgresConversationRepository struct {
	db *gorm.DB
}

// NewPostgresConversationRepository creates a new PostgresConversationRepository
func NewPostgresConversationRepository(db *gorm.DB) *PostgresConversationRepository {
	return &PostgresConversationRepository{db: db}
}

// Create inserts a new conversation into the repository
func (r *PostgresConversationRepository) Create(ctx context.Context, userID uint64, workspaceID uint64, conversation *domain.Conversation) error {
	// Verify workspace ownership
	var workspace domain.Workspace
	if err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", workspaceID, userID).
		First(&workspace).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrNotFound
		}
		return err
	}

	conversation.WorkspaceID = workspaceID
	conversation.CreatedAt = time.Now()
	conversation.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).Create(conversation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates an existing conversation in the repository
func (r *PostgresConversationRepository) Update(ctx context.Context, userID uint64, conversation *domain.Conversation) error {
	conversation.UpdatedAt = time.Now()

	result := r.db.WithContext(ctx).
		Model(conversation).
		Where("id = ? AND workspace_id IN (SELECT id FROM workspaces WHERE user_id = ? AND deleted_at IS NULL)", conversation.ID, userID).
		Updates(map[string]interface{}{
			"title":             conversation.Title,
			"provider":          conversation.Provider,
			"model":             conversation.Model,
			"system_prompt":     conversation.SystemPrompt,
			"temperature":       conversation.Temperature,
			"top_p":             conversation.TopP,
			"max_tokens":        conversation.MaxTokens,
			"streaming_enabled": conversation.StreamingEnabled,
			"updated_at":        conversation.UpdatedAt,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// Delete soft deletes a conversation from the repository
func (r *PostgresConversationRepository) Delete(ctx context.Context, userID uint64, id uint64) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.Conversation{}).
		Where("id = ? AND workspace_id IN (SELECT id FROM workspaces WHERE user_id = ? AND deleted_at IS NULL)", id, userID).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// FindByID retrieves a conversation by its ID for a specific user
func (r *PostgresConversationRepository) FindByID(ctx context.Context, userID uint64, id uint64) (*domain.Conversation, error) {
	var conversation domain.Conversation
	result := r.db.WithContext(ctx).
		Where("id = ? AND workspace_id IN (SELECT id FROM workspaces WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL", id, userID).
		First(&conversation)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &conversation, nil
}

// UpdateSummary updates the summary field of a conversation
func (r *PostgresConversationRepository) UpdateSummary(ctx context.Context, conversationID uint64, summary string) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.Conversation{}).
		Where("id = ? AND deleted_at IS NULL", conversationID).
		Updates(map[string]interface{}{
			"summary":            summary,
			"summary_updated_at": now,
			"updated_at":         now,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// FindByWorkspaceID retrieves all conversations for a specific workspace
func (r *PostgresConversationRepository) FindByWorkspaceID(ctx context.Context, userID uint64, workspaceID uint64) ([]*domain.Conversation, error) {
	var conversations []*domain.Conversation
	result := r.db.WithContext(ctx).
		Where("workspace_id = ? AND workspace_id IN (SELECT id FROM workspaces WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL", workspaceID, userID).
		Order("created_at DESC").
		Find(&conversations)
	if result.Error != nil {
		return nil, result.Error
	}
	return conversations, nil
}

// PostgresMessageRepository implements MessageRepository using PostgreSQL and GORM
type PostgresMessageRepository struct {
	db *gorm.DB
}

// NewPostgresMessageRepository creates a new PostgresMessageRepository
func NewPostgresMessageRepository(db *gorm.DB) *PostgresMessageRepository {
	return &PostgresMessageRepository{db: db}
}

// Create inserts a new message into the repository
func (r *PostgresMessageRepository) Create(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, message *domain.Message) error {
	// Verify conversation ownership via workspace
	var conversation domain.Conversation
	if err := r.db.WithContext(ctx).
		Where(`id = ? AND workspace_id = ? AND workspace_id IN (SELECT id FROM workspaces WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL`, conversationID, workspaceID, userID).
		First(&conversation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.ErrNotFound
		}
		return err
	}

	message.ConversationID = conversationID
	message.CreatedAt = time.Now()

	result := r.db.WithContext(ctx).Create(message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindByID retrieves a message by its ID for a specific user
func (r *PostgresMessageRepository) FindByID(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) (*domain.Message, error) {
	var message domain.Message
	result := r.db.WithContext(ctx).
		Where(`id = ? AND conversation_id = ? 
			AND conversation_id IN (
				SELECT c.id FROM conversations c 
				JOIN workspaces w ON c.workspace_id = w.id 
				WHERE w.user_id = ? AND w.deleted_at IS NULL AND c.deleted_at IS NULL
			) AND deleted_at IS NULL`, messageID, conversationID, userID).
		First(&message)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, result.Error
	}
	return &message, nil
}

// List retrieves all messages for a specific conversation
func (r *PostgresMessageRepository) List(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64) ([]*domain.Message, error) {
	var messages []*domain.Message
	result := r.db.WithContext(ctx).
		Where(`conversation_id = ? 
			AND conversation_id IN (
				SELECT c.id FROM conversations c 
				JOIN workspaces w ON c.workspace_id = w.id 
				WHERE w.id = ? AND w.user_id = ? AND w.deleted_at IS NULL AND c.deleted_at IS NULL
			) AND deleted_at IS NULL`, conversationID, workspaceID, userID).
		Order("created_at ASC").
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

// Delete soft deletes a message from the repository
func (r *PostgresMessageRepository) Delete(ctx context.Context, userID uint64, workspaceID uint64, conversationID uint64, messageID uint64) error {
	now := time.Now()
	result := r.db.WithContext(ctx).
		Model(&domain.Message{}).
		Where(`id = ? AND conversation_id = ? 
			AND conversation_id IN (
				SELECT c.id FROM conversations c 
				JOIN workspaces w ON c.workspace_id = w.id 
				WHERE w.id = ? AND w.user_id = ? AND w.deleted_at IS NULL AND c.deleted_at IS NULL
			) AND deleted_at IS NULL`, messageID, conversationID, workspaceID, userID).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}
