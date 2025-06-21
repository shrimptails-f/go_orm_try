-- name: GetPostWithNestedReplies :many
SELECT 
  u.id AS user_id, u.user_name as user_name,
  p.id AS post_id, p.title AS post_title,
  c.id AS comment_id, c.content AS comment_content,
  r.id AS reply_id, r.content AS reply_content
FROM users u
JOIN posts p ON p.user_id = u.id
JOIN comments c ON c.post_id = p.id
LEFT JOIN replies r ON r.comment_id = c.id
