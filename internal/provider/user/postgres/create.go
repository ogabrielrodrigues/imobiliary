package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresUserRepository) Create(ctx context.Context, user *user.User) (uuid.UUID, *response.Err) {
	row := pg.pool.QueryRow(ctx, `
		INSERT INTO "user" (id, fullname, creci_id, cellphone, email, password, avatar)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		user.ID,
		user.Fullname,
		user.CreciID,
		user.Cellphone,
		user.Email,
		user.GetPassword(),
		user.Avatar,
	)

	var id string
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: Handle specific error
	}

	return uuid.MustParse(id), nil
}
