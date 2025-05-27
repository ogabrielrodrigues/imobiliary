package provider

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresUserRepository) findByEmail(ctx context.Context, email string) (*user.User, *response.Err) {
	row := pg.pool.QueryRow(ctx, `
		SELECT
			id,
			password
		FROM "user"
		WHERE email = $1`, email)

	var found user.User
	var password string
	if err := row.Scan(&found.ID, &password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.NewErr(http.StatusNotFound, user.ERR_USER_NOT_FOUND_OR_NOT_EXISTS)
		}

		return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	found.SetPassword(password)

	return &found, nil
}
