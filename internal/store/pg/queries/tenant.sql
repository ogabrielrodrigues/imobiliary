-- name: GetTenant :one
SELECT
  *
FROM tenant
WHERE id = $1;

---------------------------------

-- name: GetTenants :many
SELECT
  *
FROM tenant;

---------------------------------

-- name: InsertTenant :one
INSERT INTO tenant
  (fullname, rg, cpf, occupation, marital_status) VALUES
  ($1, $2, $3, $4, $5)
RETURNING id;

---------------------------------

-- name: UpdateTenant :exec
UPDATE tenant SET
  fullname = $1,
  rg = $2,
  cpf = $3,
  occupation = $4,
  marital_status = $5
WHERE id = $6;

---------------------------------

-- name: DeleteTenant :exec
DELETE FROM tenant
WHERE id = $1;