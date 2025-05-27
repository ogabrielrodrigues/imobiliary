package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresOwnerRepository) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	_, err := pg.pool.Begin(ctx)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	owner_query := `SELECT id FROM "owner" WHERE id = $1`
	row := pg.pool.QueryRow(ctx, owner_query, owner_id)

	if err := row.Scan(&owner_id); err != nil {
		if err == pgx.ErrNoRows {
			return response.NewErr(http.StatusNotFound, owner.ERR_OWNER_NOT_FOUND_OR_NOT_EXISTS)
		}

		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	property_query := `SELECT id FROM "property" WHERE id = $1`
	row = pg.pool.QueryRow(ctx, property_query, property_id)

	if err := row.Scan(&property_id); err != nil {
		if err == pgx.ErrNoRows {
			return response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
		}

		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	query := `
		UPDATE "property"
			SET owner_id = $1
		WHERE id = $2 AND manager_id = $3
	`

	_, err = pg.pool.Exec(ctx, query, owner_id, property_id, user_id)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}
