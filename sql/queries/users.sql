-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING id, created_at, updated_at, email, hashed_password;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, hashed_password
FROM users
WHERE email = $1; 

-- name: DeleteUsers :exec
DELETE FROM users;