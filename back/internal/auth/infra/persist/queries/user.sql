-- name: CreateUser :exec
INSERT INTO account.users (id, first_name, last_name, email, password) 
VALUES ($1, $2, $3, $4, $5);

-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password
FROM account.users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, first_name, last_name, email, password
FROM account.users
WHERE id = $1;