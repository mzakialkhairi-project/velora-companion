-- Remove conversation summary fields
ALTER TABLE conversations DROP COLUMN IF EXISTS summary;
ALTER TABLE conversations DROP COLUMN IF EXISTS summary_updated_at;
