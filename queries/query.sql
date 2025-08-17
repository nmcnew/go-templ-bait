-- name: GetUser :one
SELECT u.*
FROM user u
WHERE u.id = ?;
