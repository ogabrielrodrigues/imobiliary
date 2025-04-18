package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresUserRepository) Authenticate(ctx context.Context, email, password string) (*user.User, *response.Err) {
	found, err := pg.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !found.ComparePwd(password) {
		return found, response.NewErr(http.StatusUnauthorized, user.ERR_PASSWORD_DONT_MATCH)
	}

	return found, nil
}

func (pg *PostgresUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	id := ctx.Value("user_id").(string)
	user_id, r_err := uuid.Parse(id)
	if r_err != nil {
		return response.NewErr(http.StatusBadRequest, user.ERR_UUID_INVALID)
	}

	_, err := pg.pool.Exec(ctx, `
		UPDATE "user"
		SET avatar = $1
		WHERE id = $2`,
		avatar_url,
		user_id,
	)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: Handle specific error
	}

	return nil
}
