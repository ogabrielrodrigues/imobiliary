package provider

import (
	"context"
	"errors"
	"net/http"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (pg *PostgresOwnerRepository) FindByID(ctx context.Context, owner_id uuid.UUID) (*owner.DTO, *response.Err) {
	query := `
		SELECT
			ow.id,
			ow.fullname,
			ow.cpf,
			ow.rg,
			ow.email,
			ow.cellphone,
			ow.occupation,
			ow.marital_status,
			ow.manager_id,
			ad.street,
			ad.number,
			ad.complement,
			ad.neighborhood,
			ad.city,
			ad.state,
			ad.zip_code,
			ad.mini_address
		FROM "owner" ow
		JOIN "address" ad
		ON ow.address_id = ad.id
		WHERE ow.id = $1`

	row := pg.pool.QueryRow(ctx, query, owner_id)

	var found owner.DTO
	if err := row.Scan(
		&found.ID,
		&found.Fullname,
		&found.CPF,
		&found.RG,
		&found.Email,
		&found.Cellphone,
		&found.Occupation,
		&found.MaritalStatus,
		&found.ManagerID,
		&found.Address.Street,
		&found.Address.Number,
		&found.Address.Complement,
		&found.Address.Neighborhood,
		&found.Address.City,
		&found.Address.State,
		&found.Address.ZipCode,
		&found.Address.MiniAddress,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, response.NewErr(http.StatusNotFound, owner.ERR_OWNER_NOT_FOUND_OR_NOT_EXISTS)
		}

		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	return &found, nil
}
