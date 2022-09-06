BEGIN;

-- DROP TABLE IF EXISTS tags_videos;
-- DROP TABLE IF EXISTS tags;
-- DROP TABLE IF EXISTS videos;

DROP INDEX IF EXISTS posts_audit_idx_id;
DROP TABLE IF EXISTS posts_audit;
DROP TABLE IF EXISTS posts;
DROP FUNCTION IF EXISTS update_posts_trigger_audit;

COMMIT;