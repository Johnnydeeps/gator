-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
select *
from users
where name = $1;

-- name: ResetDB :exec
delete from users;

-- name: GetAllUsers :many
select *
from users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;
