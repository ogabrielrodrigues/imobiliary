package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresPropertyRepository) FindByID(ctx context.Context, id uuid.UUID) (*property.DTO, *response.Err) {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, lib.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	query := `
	SELECT
		pr.id,
		pr.status,
		pr.kind,
		pr.water_id,
		pr.energy_id,
		pr.owner_id,
		ad.street,
		ad.number,
		ad.complement,
		ad.neighborhood,
		ad.city,
		ad.state,
		ad.zip_code,
		ad.mini_address
	FROM "property" pr
	JOIN "address" ad
	ON pr.address_id = ad.id
	WHERE
		pr.id = $1
	AND
		pr.manager_id = $2`

	row := pg.pool.QueryRow(ctx, query, id, user_id)

	var owner_id *string
	var p property.DTO
	if err := row.Scan(
		&p.ID,
		&p.Status,
		&p.Kind,
		&p.WaterID,
		&p.EnergyID,
		&owner_id,
		&p.Address.Street,
		&p.Address.Number,
		&p.Address.Complement,
		&p.Address.Neighborhood,
		&p.Address.City,
		&p.Address.State,
		&p.Address.ZipCode,
		&p.Address.MiniAddress,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
		}

		return nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	if owner_id != nil {
		p.OwnerID = *owner_id
	}

	return &p, nil
}
