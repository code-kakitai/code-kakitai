-- name: UserFindById :one
SELECT
   *
FROM
   users
WHERE
   id = ?;

-- name: CreateUser :exec
INSERT INTO
   users (id, email, password, created_at, updated_at)
VALUES
   (
      sqlc.arg(id),
      sqlc.arg(email),
      sqlc.arg(password),
      NOW(),
      NOW()
   );

-- name: UpdateUser :exec
UPDATE
   users
SET
   email = sqlc.arg(email),
   phone_number = sqlc.arg(phone_number),
   name = sqlc.arg(name),
   postal_code = sqlc.arg(postal_code),
   prefecture = sqlc.arg(prefecture),
   city = sqlc.arg(city),
   address_extra = sqlc.arg(address_extra),
   updated_at = NOW();