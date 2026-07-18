-- Create conversations table
CREATE TABLE conversations (
    id BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT NOT NULL,
    title VARCHAR(100),
    provider VARCHAR(50),
    model VARCHAR(50),
    system_prompt TEXT,
    temperature DECIMAL,
    max_tokens INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_conversations_workspace_id ON conversations(workspace_id);
CREATE INDEX idx_conversations_deleted_at ON conversations(deleted_at);

-- Add foreign key constraint
ALTER TABLE conversations ADD CONSTRAINT fk_conversations_workspace_id FOREIGN KEY (workspace_id) REFERENCES workspaces(id);
