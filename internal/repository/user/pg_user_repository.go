package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_CONNECT_TO_DB = "failed to connect to database"
)

type PostgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresUserRepository(pool *pgxpool.Pool) (*PostgresUserRepository, *response.Err) {
	if err := pool.Ping(context.Background()); err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, ERR_CONNECT_TO_DB)
	}

	return &PostgresUserRepository{pool}, nil
}

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

		return nil, response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: Handle specific error
	}

	found.SetPassword(*pwd)

	return &found, nil
}

func (pg *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, *response.Err) {
	row := pg.pool.QueryRow(ctx, `
		SELECT
			*
		FROM "user"
		WHERE email = $1`, email)

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

		return nil, response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: Handle specific error
	}

	found.SetPassword(*pwd)

	return &found, nil
}

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

func (pg *PostgresUserRepository) ChangeAvatar(ctx context.Context, id uuid.UUID, avatar_url string) *response.Err {
	_, err := pg.pool.Exec(ctx, `
		UPDATE "user"
		SET avatar = $1
		WHERE id = $2`,
		avatar_url,
		id.String(),
	)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, err.Error()) // TODO: Handle specific error
	}

	return nil
}
