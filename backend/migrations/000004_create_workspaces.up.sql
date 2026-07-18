-- Create workspaces table
CREATE TABLE workspaces (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_workspaces_user_id ON workspaces(user_id);
CREATE INDEX idx_workspaces_name ON workspaces(name);
CREATE INDEX idx_workspaces_deleted_at ON workspaces(deleted_at);

-- Add foreign key constraint
ALTER TABLE workspaces ADD CONSTRAINT fk_workspaces_user_id FOREIGN KEY (user_id) REFERENCES users(id);
