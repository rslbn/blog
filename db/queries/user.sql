-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
-- name: ListUsers :many
SELECT * FROM users;
-- name: UserExistByUsername :one
SELECT COUNT(*) > 0 AS is_exist FROM users WHERE username = $1;
-- name: UserExistsByEmail :one
SELECT COUNT(*) > 0 AS is_exist FROM users WHERE email = $1;
-- name: InsertUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3) RETURNING *;
-- name: UpdateUser :one
UPDATE users SET username = $1, email = $2, password = $3
WHERE username = $1 RETURNING *;
-- name: DeleteUserByID :one
DELETE FROM users WHERE user_id = $1 RETURNING *;
-- name: DeleteByUsername :one
DELETE FROM users WHERE username = $1 RETURNING *;