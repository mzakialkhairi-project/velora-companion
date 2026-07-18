-- Add conversation AI settings
ALTER TABLE conversations ADD COLUMN top_p DECIMAL(3,2) DEFAULT 1.0;
ALTER TABLE conversations ADD COLUMN streaming_enabled BOOLEAN DEFAULT true;

-- Add comments
COMMENT ON COLUMN conversations.top_p IS 'Controls diversity via nucleus sampling (0.0-1.0)';
COMMENT ON COLUMN conversations.streaming_enabled IS 'Enable streaming responses for this conversation';
