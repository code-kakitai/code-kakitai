-- name: InsertOrderProduct :exec
INSERT INTO
  order_products (
      id,
      order_id,
      product_id,
      price,
      quantity
  )
VALUES
  (
      sqlc.arg(id),
      sqlc.arg(order_id),
      sqlc.arg(product_id),
      sqlc.arg(price),
      sqlc.arg(quantity)
  );

-- name: OrderProductFindById :one
SELECT
  *
FROM
  order_products
WHERE
  id = ?;

