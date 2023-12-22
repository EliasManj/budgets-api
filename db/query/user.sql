-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
order by username
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password, email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;