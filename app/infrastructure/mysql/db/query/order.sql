-- name: InsertOrder :exec
INSERT INTO
  orders (
      id,
      user_id,
      total_amount,
      ordered_at
  )
VALUES
  (
      sqlc.arg(id),
      sqlc.arg(user_id),
      sqlc.arg(total_amount),
      sqlc.arg(ordered_at)
  );

-- name: OrderFindById :one
SELECT
  *
FROM
  orders
WHERE
  id = ?;