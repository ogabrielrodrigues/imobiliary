-- -- name: GetAddress :one
-- SELECT
--   *
-- FROM address
-- WHERE id = $1;

-- ---------------------------------

-- -- name: GetAddresses :many
-- SELECT
--   *
-- FROM address;

-- ---------------------------------

-- -- name: InsertAddress :one
-- INSERT INTO address
--   (street, number, district, city, cep, kind, tenant_id, owner_id) VALUES
--   ($1, $2, $3, $4, $5, $6, $7, $8)
-- RETURNING id;

-- ---------------------------------

-- -- name: UpdateAddress :exec
-- UPDATE
--   address
-- SET
--   street = $1,
--   number = $2,
--   district = $3,
--   city = $4,
--   cep = $5,
--   kind = $6,
--   tenant_id = $7,
--   owner_id = $8
-- WHERE id = $9;

-- ---------------------------------

-- -- name: DeleteAddress :exec
-- DELETE FROM
--   address
-- WHERE id = $1;

