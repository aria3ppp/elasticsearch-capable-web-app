SELECT  id,
        title,
        body,
        modified_at,
        is_deleted
FROM posts
WHERE modified_at > :sql_last_value
ORDER BY modified_at ASC;