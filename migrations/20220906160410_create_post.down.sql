BEGIN;

DROP TABLE IF EXISTS posts_audit;
DROP TABLE IF EXISTS posts;
DROP FUNCTION IF EXISTS update_posts_trigger_audit;

COMMIT;