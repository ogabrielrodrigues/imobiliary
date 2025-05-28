package postgres

import (
	"context"
	"database/sql"
	"errors"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"
	"imobiliary/internal/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresManagerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresManagerRepository(pool *pgxpool.Pool) *PostgresManagerRepository {
	return &PostgresManagerRepository{db: pool}
}

func (mr *PostgresManagerRepository) FindByID(ctx context.Context, managerID uuid.UUID) (*manager.Manager, *response.Err) {
	row := mr.db.QueryRow(ctx, `SELECT * FROM "user" WHERE id = $1`, managerID)

	var found manager.Manager
	if err := row.Scan(
		&found.ID,
		&found.Fullname,
		&found.Phone,
		&found.Email,
		&found.Password,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.NewErr(http.StatusNotFound, err.Error()) //  TODO: refactor error
		}

		return nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return &found, nil
}

func (mr *PostgresManagerRepository) Create(ctx context.Context, manager *manager.Manager) *response.Err {
	_, err := mr.db.Exec(ctx, `
		INSERT INTO "user" (id, fullname, phone, email, password)
		VALUES ($1, $2, $3, $4, $5)`,
		manager.ID,
		manager.Fullname,
		manager.Phone,
		manager.Email,
		manager.Password,
	)

	if err != nil {
		if IsUniqueConstraint(err) {
			return response.NewErr(http.StatusConflict, err.Error()) //  TODO: refactor error
		}

		return response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return nil
}

func (mr *PostgresManagerRepository) Authenticate(ctx context.Context, email *types.Email, password string) (uuid.UUID, *response.Err) {
	row := mr.db.QueryRow(ctx, `SELECT id, password FROM "user" WHERE email = $1`, email)

	var found manager.Manager
	if err := row.Scan(&found.ID, &found.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, response.NewErr(http.StatusNotFound, err.Error()) //  TODO: refactor error
		}

		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	if !found.ComparePassword(password) {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, "") //  TODO: refactor error
	}

	return found.ID, nil
}
