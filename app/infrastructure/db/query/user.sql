-- name: FetchByUserId :one
SELECT
   id,
   email,
   phone_number,
   name,
   postal_code,
   prefecture,
   city,
   address_extra,
   created_at,
   updated_at
FROM
   users
WHERE id = ?;