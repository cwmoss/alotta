-- name: GetBlog :one
SELECT
    *
FROM
    blogs
WHERE
    id = ?
LIMIT
    1;

-- name: ListBlogs :many
SELECT
    *
FROM
    blogs
ORDER BY
    created_at DESC;

-- name: CreateBlog :one
INSERT INTO
    blogs (title, content, created_at)
VALUES
    (?, ?, DateTime ('now')) RETURNING *;

-- name: UpdateBlog :exec
UPDATE blogs
set
    title = ?,
    content = ?
WHERE
    id = ?;

-- name: DeleteAuthor :exec
DELETE FROM blogs
WHERE
    id = ?;