-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CountUsers :one
SELECT count(*) FROM users;

-- name: CreateUser :exec
INSERT INTO users (username, password, email) VALUES ($1, $2, $3);

-- name: UpdateUser :exec
UPDATE users SET username = $2, password = $3, email = $4 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;