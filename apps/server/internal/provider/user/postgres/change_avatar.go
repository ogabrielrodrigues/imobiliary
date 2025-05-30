package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/middleware"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (pg *PostgresUserRepository) ChangeAvatar(ctx context.Context, avatar_url string) *response.Err {
	id := ctx.Value(middleware.UserIDKey).(string)
	user_id, r_err := uuid.Parse(id)
	if r_err != nil {
		return response.NewErr(http.StatusBadRequest, response.ERR_INVALID_UUID)
	}

	_, err := pg.pool.Exec(ctx, `
		UPDATE "user"
		SET avatar = $1
		WHERE id = $2`,
		avatar_url,
		user_id,
	)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}
