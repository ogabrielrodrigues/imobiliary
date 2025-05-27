package provider

import (
	"context"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (pg *PostgresUserRepository) ListAll(ctx context.Context) ([]user.DTO, *response.Err) {
	rows, err := pg.pool.Query(ctx, `SELECT * FROM "user"`)
	if err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}
	defer rows.Close()

	var users []user.DTO
	for rows.Next() {
		var u user.DTO
		var pwd *string
		if err := rows.Scan(
			&u.ID,
			&u.Fullname,
			&u.CreciID,
			&u.Cellphone,
			&u.Email,
			&pwd,
			&u.Avatar,
		); err != nil {
			return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return users, nil
}
