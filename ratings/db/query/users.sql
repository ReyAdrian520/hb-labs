-- name: CreateUser :one
INSERT INTO users(
    username,
    password,
    email
)   VALUES(
        $1, $2, $3 
) RETURNING *;

-- name: GetUsersDetails :one
SELECT * FROM users
WHERE  user_id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM USERS WHERE 'title' = $1;