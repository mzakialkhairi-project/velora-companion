-- Remove conversation AI settings
ALTER TABLE conversations DROP COLUMN IF EXISTS top_p;
ALTER TABLE conversations DROP COLUMN IF EXISTS streaming_enabled;
