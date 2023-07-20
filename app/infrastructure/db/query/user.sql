-- name: UserFindById :one
SELECT
   *
FROM
   users
WHERE id = ?;