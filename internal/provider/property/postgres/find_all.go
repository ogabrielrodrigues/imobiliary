package provider

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresPropertyRepository) FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err) {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	query := `
		SELECT
			pr.id,
			pr.status,
			pr.kind,
			pr.water_id,
			pr.energy_id,
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
		AND
			pr.manager_id = $1`

	rows, err := pg.pool.Query(ctx, query, user_id)
	if err != nil {
		if pgx.ErrNoRows == err {
			return nil, response.NewErr(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}

		return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}
	defer rows.Close()

	var properties []property.DTO
	for rows.Next() {
		var p property.DTO
		if err := rows.Scan(
			&p.ID,
			&p.Status,
			&p.Kind,
			&p.WaterID,
			&p.EnergyID,
			&p.Address.Street,
			&p.Address.Number,
			&p.Address.Complement,
			&p.Address.Neighborhood,
			&p.Address.City,
			&p.Address.State,
			&p.Address.ZipCode,
			&p.Address.MiniAddress,
		); err != nil {
			return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
		}

		properties = append(properties, p)
	}

	return properties, nil
}
