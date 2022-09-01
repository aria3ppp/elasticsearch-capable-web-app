BEGIN;

-- Posts table to write to
CREATE TABLE IF NOT EXISTS posts_store (
    id SERIAL,
    title VARCHAR(150) NOT NULL,
    body TEXT NOT NULL,
    modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY (id, modified_at)
);

-- Posts logs for Logstash to read from
CREATE OR REPLACE VIEW posts_log AS
    SELECT ps.id, ps.title, ps.body, ps.modified_at, ps.is_deleted
    FROM posts_store ps INNER JOIN
        (
            SELECT id, MAX(modified_at) AS modified_at
            FROM posts_store
            GROUP BY id 
        ) g
    ON ps.id = g.id AND ps.modified_at = g.modified_at;

-- Posts view to read from
CREATE OR REPLACE VIEW posts AS
    SELECT id, title, body, modified_at, is_deleted
    FROM posts_log
    WHERE NOT is_deleted;

COMMIT;