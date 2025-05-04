package provider

import (
	"context"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresPropertyRepository) Create(ctx context.Context, property *property.Property) *response.Err {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		logger.Error(err.Error())
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	address_query := `
		INSERT INTO "address" (street, "number", complement, neighborhood, city, state, zip_code, full_address, mini_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	row := tx.QueryRow(ctx, address_query,
		property.Address.Street,
		property.Address.Number,
		property.Address.Complement,
		property.Address.Neighborhood,
		property.Address.City,
		property.Address.State,
		property.Address.ZipCode,
		property.Address.FullAddress,
		property.Address.MiniAddress,
	)

	var address_id string
	if err := row.Scan(&address_id); err != nil {
		logger.Error(err.Error())
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	manager_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	logger.Info(manager_id)

	property_query := `
		INSERT INTO "property" (id, status, kind, water_id, energy_id, manager_id, address_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = tx.Exec(ctx, property_query,
		property.ID,
		property.Status,
		property.Kind,
		property.WaterID,
		property.EnergyID,
		manager_id,
		address_id,
	)

	if err != nil {
		logger.Error(err.Error())
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error(err.Error())
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}
