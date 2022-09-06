SELECT  id,
        title,
        body,
        contributed_at,
        deleted
FROM posts
WHERE contributed_at > :sql_last_value
ORDER BY contributed_at ASC;