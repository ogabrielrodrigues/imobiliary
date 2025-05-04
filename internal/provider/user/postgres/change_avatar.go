package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	id := ctx.Value(middleware.UserIDKey).(string)
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
		return response.NewErr(http.StatusInternalServerError, user.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}
