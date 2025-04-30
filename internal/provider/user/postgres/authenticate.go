package provider

import (
	"context"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresUserRepository) Authenticate(ctx context.Context, email, password string) (*user.User, *response.Err) {
	found, err := pg.findByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !found.ComparePwd(password) {
		return found, response.NewErr(http.StatusUnauthorized, user.ERR_PASSWORD_DONT_MATCH)
	}

	return found, nil
}
