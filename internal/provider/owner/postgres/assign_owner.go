package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresOwnerRepository) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	query := `
		UPDATE "property"
			SET owner_id = $1
		WHERE id = $2 AND manager_id = $3
	`

	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	_, err := pg.pool.Exec(ctx, query, owner_id, property_id, user_id)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error())
	}

	return nil
}
