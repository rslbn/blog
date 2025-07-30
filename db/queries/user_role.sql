-- name: GetUserRoleByUserID :many
SELECT * FROM user_role WHERE user_id = $1;
-- name: GetUserRoleByRoleID :many
SELECT * FROM user_role WHERE role_id = $1;
-- name: ListUserRoles :many
SELECT * FROM user_role;
-- name: InsertUserRole :one
INSERT INTO user_role (user_id, role_id)
VALUES ($1, $2) RETURNING *;
-- name: DeleteUserRoleByUserID :exec
DELETE FROM user_role WHERE user_id = $1 RETURNING *;
-- name: DeleteUserRoleByUserIDAndRoleID :exec
DELETE FROM user_role WHERE user_id = $1 AND role_id = $2 RETURNING *;