-- Add conversation summary fields
ALTER TABLE conversations ADD COLUMN summary TEXT;
ALTER TABLE conversations ADD COLUMN summary_updated_at TIMESTAMP;
