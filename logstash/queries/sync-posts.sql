SELECT  id,
        title,
        body,
        contributed_at,
        deleted
FROM posts
WHERE contributed_at > ?
ORDER BY contributed_at ASC;