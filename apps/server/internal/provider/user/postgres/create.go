package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"
	"imobiliary/internal/store"

	"github.com/google/uuid"
)

func (pg *PostgresUserRepository) Create(ctx context.Context, usr *user.User) (uuid.UUID, *response.Err) {
	row := pg.pool.QueryRow(ctx, `
		INSERT INTO "user" (id, fullname, creci_id, cellphone, email, password, avatar)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		usr.ID,
		usr.Fullname,
		usr.CreciID,
		usr.Cellphone,
		usr.Email,
		usr.GetPassword(),
		usr.Avatar,
	)

	var id string
	if err := row.Scan(&id); err != nil {
		if store.IsUniqueConstraint(err) {
			return uuid.Nil, response.NewErr(http.StatusConflict, user.ERR_USER_ALREADY_EXISTS)
		}

		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return uuid.MustParse(id), nil
}
