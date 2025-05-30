package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (pg *PostgresUserRepository) Authenticate(ctx context.Context, dto *user.AuthDTO) (uuid.UUID, *response.Err) {
	found, err := pg.findByEmail(ctx, dto.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if !found.ComparePwd(dto.Password) {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, user.ERR_PASSWORD_DONT_MATCH)
	}

	return found.ID, nil
}
