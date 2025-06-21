-- name: GetUserPostsWithComments :many
SELECT
  users.id AS user_id,
  users.user_name,
  posts.id AS post_id,
  posts.title,
  comments.id AS comment_id,
  comments.content
FROM users
JOIN posts ON users.id = posts.user_id
LEFT JOIN comments ON posts.id = comments.post_id
WHERE users.id = $1
ORDER BY posts.id, comments.id;
