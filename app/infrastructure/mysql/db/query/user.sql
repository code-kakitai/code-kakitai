-- name: UserFindById :one
SELECT
   *
FROM
   users
WHERE
   id = ?;

-- name: UpsertUser :exec
INSERT INTO
   users (
      id,
      email,
      firebaseUid,
      phone_number,
      first_name,
      last_name,
      postal_code,
      prefecture,
      city,
      address_extra,
      created_at,
      updated_at
   )
VALUES
   (
      sqlc.arg(id),
      sqlc.arg(email),
      sqlc.arg(firebaseUid),
      sqlc.arg(phone_number),
      sqlc.arg(first_name),
      sqlc.arg(last_name),
      sqlc.arg(postal_code),
      sqlc.arg(prefecture),
      sqlc.arg(city),
      sqlc.arg(address_extra),
      NOW(),
      NOW()
   ) ON DUPLICATE KEY
UPDATE
   email = sqlc.arg(email),
   firebaseUid = sqlc.arg(firebaseUid),
   phone_number = sqlc.arg(phone_number),
   first_name = sqlc.arg(first_name),
   last_name = sqlc.arg(last_name),
   postal_code = sqlc.arg(postal_code),
   prefecture = sqlc.arg(prefecture),
   city = sqlc.arg(city),
   address_extra = sqlc.arg(address_extra),
   updated_at = NOW();

-- name: UserFindAll :many
SELECT
   *
FROM
   users;