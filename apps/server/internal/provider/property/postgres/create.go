package provider

import (
	"context"
	"net/http"

	"imobiliary/config/logger"
	"imobiliary/internal/entity/property"
	jwt "imobiliary/internal/lib"
	"imobiliary/internal/middleware"
	"imobiliary/internal/response"
	"imobiliary/internal/store"
)

func (pg *PostgresPropertyRepository) Create(ctx context.Context, dto *property.Property) *response.Err {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	address_query := `
		INSERT INTO "address" (street, "number", complement, neighborhood, city, state, zip_code, full_address, mini_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	row := tx.QueryRow(ctx, address_query,
		dto.Address.Street,
		dto.Address.Number,
		dto.Address.Complement,
		dto.Address.Neighborhood,
		dto.Address.City,
		dto.Address.State,
		dto.Address.ZipCode,
		dto.Address.FullAddress,
		dto.Address.MiniAddress,
	)

	var address_id string
	if err := row.Scan(&address_id); err != nil {
		tx.Rollback(ctx)

		if store.IsUniqueConstraint(err) {
			return response.NewErr(http.StatusConflict, property.ERR_PROPERTY_ALREADY_EXISTS)
		}

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
		dto.ID,
		dto.Status,
		dto.Kind,
		dto.WaterID,
		dto.EnergyID,
		manager_id,
		address_id,
	)

	if err != nil {
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}
