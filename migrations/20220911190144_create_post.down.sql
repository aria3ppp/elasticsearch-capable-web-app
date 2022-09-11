BEGIN;

DROP FUNCTION IF EXISTS posts_function_triggers_on_update;
DROP TABLE IF EXISTS posts_audit;
DROP TABLE IF EXISTS posts;

COMMIT;