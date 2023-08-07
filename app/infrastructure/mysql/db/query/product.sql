-- name: ProductFindById :one
SELECT
   *
FROM
   products
WHERE
   id = ?;

-- name: ProductFindByIds :many
SELECT
   *
FROM
   products
WHERE
   id IN (sqlc.slice('ids'));

-- name: UpsertProduct :exec
INSERT INTO products (
   id,
   owner_id,
   name,
   description,
   price,
   stock
) VALUES (
   sqlc.arg(id),
   sqlc.arg(owner_id),
   sqlc.arg(name),
   sqlc.arg(description),
   sqlc.arg(price),
   sqlc.arg(stock)
) ON DUPLICATE KEY UPDATE
   owner_id = sqlc.arg(owner_id),
   name = sqlc.arg(name),
   description = sqlc.arg(description),
   price = sqlc.arg(price),
   stock = sqlc.arg(stock)