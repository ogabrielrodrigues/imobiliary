package provider

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*user.User, *response.Err) {
	row := pg.pool.QueryRow(ctx, `
		SELECT
			*
		FROM "user"
		WHERE id = $1`, id)

	var pwd *string
	var found user.User
	if err := row.Scan(
		&found.ID,
		&found.Fullname,
		&found.CreciID,
		&found.Cellphone,
		&found.Email,
		&pwd,
		&found.Avatar,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
		}

		return nil, response.NewErr(http.StatusInternalServerError, user.ERR_INTERNAL_SERVER_ERROR) // TODO: Handle specific error
	}

	found.SetPassword(*pwd)

	return &found, nil
}
