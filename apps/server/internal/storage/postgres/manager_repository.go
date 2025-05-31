package postgres

import (
	"context"
	"database/sql"
	"errors"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresManagerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresManagerRepository(pool *pgxpool.Pool) *PostgresManagerRepository {
	return &PostgresManagerRepository{db: pool}
}

func (mr *PostgresManagerRepository) FindByID(ctx context.Context, managerID uuid.UUID) (*manager.Manager, *httperr.HttpError) {
	row := mr.db.QueryRow(ctx, `SELECT * FROM "manager" WHERE id = $1`, managerID)

	var found manager.Manager
	if err := row.Scan(
		&found.ID,
		&found.Fullname,
		&found.Phone.Number,
		&found.Email.Address,
		&found.Password,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httperr.NewNotFoundError(ctx, "manager not found or not exists")
		}

		return nil, httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	return &found, nil
}

func (mr *PostgresManagerRepository) Create(ctx context.Context, manager *manager.Manager) *httperr.HttpError {
	_, err := mr.db.Exec(ctx, `
		INSERT INTO "manager" (id, fullname, phone, email, password)
		VALUES ($1, $2, $3, $4, $5)`,
		manager.ID,
		manager.Fullname,
		manager.Phone.Value(),
		manager.Email.Value(),
		manager.Password,
	)

	if err != nil {
		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "manager already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	return nil
}

func (mr *PostgresManagerRepository) Authenticate(ctx context.Context, email *types.Email, password string) (uuid.UUID, *httperr.HttpError) {
	row := mr.db.QueryRow(ctx, `SELECT id, password FROM "manager" WHERE email = $1`, email.Value())

	var found manager.Manager
	if err := row.Scan(&found.ID, &found.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, httperr.NewNotFoundError(ctx, "manager not found or not exists")
		}

		return uuid.Nil, httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	if !found.ComparePassword(password) {
		return uuid.Nil, httperr.NewUnauthorizedError(ctx, "password don't match")
	}

	return found.ID, nil
}
