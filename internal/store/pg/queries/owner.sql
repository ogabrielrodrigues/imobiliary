-- name: GetOwner :one
SELECT
  *
FROM owner
WHERE id = $1;

---------------------------------

-- name: GetOwners :many
SELECT
  *
FROM owner;

---------------------------------

-- name: InsertOwner :one
INSERT INTO owner
  (fullname, rg, cpf, occupation, marital_status) VALUES
  ($1, $2, $3, $4, $5)
RETURNING id;

---------------------------------

-- name: UpdateOwner :exec
UPDATE
  owner
SET
  fullname = $1,
  rg = $2,
  cpf = $3,
  occupation = $4,
  marital_status = $5
WHERE id = $6;

---------------------------------

-- name: DeleteOwner :exec
DELETE FROM
  owner
WHERE id = $1;