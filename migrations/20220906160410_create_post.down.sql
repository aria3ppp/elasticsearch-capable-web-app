BEGIN;

DROP FUNCTION IF EXISTS update_posts_trigger_audit;
DROP TABLE IF EXISTS posts_audit;
DROP TABLE IF EXISTS posts;

COMMIT;