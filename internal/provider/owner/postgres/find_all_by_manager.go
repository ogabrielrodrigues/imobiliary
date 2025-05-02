package provider

import (
	"context"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresOwnerRepository) FindAllByManagerID(ctx context.Context) ([]owner.DTO, *response.Err) {
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
		WHERE ow.manager_id = $1`

	manager_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	rows, err := pg.pool.Query(ctx, query, manager_id)
	if err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var owners []owner.DTO
	for rows.Next() {
		var o owner.DTO
		if err := rows.Scan(
			&o.ID,
			&o.Fullname,
			&o.CPF,
			&o.RG,
			&o.Email,
			&o.Cellphone,
			&o.Occupation,
			&o.MaritalStatus,
			&o.ManagerID,
			&o.Address.Street,
			&o.Address.Number,
			&o.Address.Complement,
			&o.Address.Neighborhood,
			&o.Address.City,
			&o.Address.State,
			&o.Address.ZipCode,
			&o.Address.MiniAddress,
		); err != nil {
			return nil, response.NewErr(http.StatusInternalServerError, err.Error())
		}

		owners = append(owners, o)
	}

	return owners, nil
}
