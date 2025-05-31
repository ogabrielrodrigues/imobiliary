package postgres

import (
	"context"
	"database/sql"
	"errors"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/owner"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOwnerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresOwnerRepository(pool *pgxpool.Pool) *PostgresOwnerRepository {
	return &PostgresOwnerRepository{db: pool}
}

func (mr *PostgresOwnerRepository) FindByID(ctx context.Context, ownerID uuid.UUID, managerID uuid.UUID) (*owner.Owner, *httperr.HttpError) {
	row := mr.db.QueryRow(ctx, `
	SELECT
		ow.id,
		ow.manager_id,
		ow.fullname,
		ow.cpf,
		ow.rg,
		ow.phone,
		ow.email,
		ow.occupation,
		ow.marital_status,
		ad.full_address,
		ad.street,
		ad.number,
		ad.complement,
		ad.neighborhood,
		ad.city,
		ad.state,
		ad.zip_code
	FROM "owner" ow
	INNER JOIN "address" ad ON ow.address_id = ad.id
	WHERE ow.id = $1 AND ow.manager_id = $2`, ownerID, managerID)

	var found owner.Owner
	if err := row.Scan(
		&found.ID,
		&found.ManagerID,
		&found.Fullname,
		&found.CPF.Cpf,
		&found.RG.Rg,
		&found.Phone.Number,
		&found.Email.Address,
		&found.Occupation,
		&found.MaritalStatus,
		&found.Address.FullAddress,
		&found.Address.Street,
		&found.Address.Number,
		&found.Address.Complement,
		&found.Address.Neighborhood,
		&found.Address.City,
		&found.Address.State,
		&found.Address.ZipCode,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httperr.NewNotFoundError(ctx, "owner not found or not exists")
		}

		return nil, httperr.NewInternalServerError(ctx, err.Error())
	}

	return &found, nil
}

func (mr *PostgresOwnerRepository) Create(ctx context.Context, owner *owner.Owner, managerID uuid.UUID) *httperr.HttpError {
	tx, err := mr.db.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error creating owner address")
	}

	row := tx.QueryRow(ctx, `
		INSERT INTO "address"
		(full_address, street, number, complement, neighborhood, city, state, zip_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`,
		owner.Address.FullAddress,
		owner.Address.Street,
		owner.Address.Number,
		owner.Address.Complement,
		owner.Address.Neighborhood,
		owner.Address.City,
		owner.Address.State,
		owner.Address.ZipCode,
	)

	var addressID string
	if err := row.Scan(&addressID); err != nil {
		tx.Rollback(ctx)

		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "owner address already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO "owner"
		(id, manager_id, address_id, fullname, cpf, rg, phone, email, occupation, marital_status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		owner.ID,
		managerID,
		addressID,
		owner.Fullname,
		owner.CPF.Value(),
		owner.RG.Value(),
		owner.Phone.Value(),
		owner.Email.Value(),
		owner.Occupation,
		owner.MaritalStatus,
	)

	if err != nil {
		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "owner already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error committing transaction")
	}

	return nil
}

func (mr *PostgresOwnerRepository) FindAll(ctx context.Context, managerID uuid.UUID) ([]owner.Owner, *httperr.HttpError) {
	rows, err := mr.db.Query(ctx, `
	SELECT
		ow.id,
		ow.manager_id,
		ow.fullname,
		ow.cpf,
		ow.rg,
		ow.phone,
		ow.email,
		ow.occupation,
		ow.marital_status,
		ad.full_address,
		ad.street,
		ad.number,
		ad.complement,
		ad.neighborhood,
		ad.city,
		ad.state,
		ad.zip_code
	FROM "owner" ow
	INNER JOIN "address" ad ON ow.address_id = ad.id
	WHERE ow.manager_id = $1`, managerID)

	if err != nil {
		if IsErrNoRows(err) {
			return []owner.Owner{}, nil
		}

		return nil, httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	var found owner.Owner
	var owners []owner.Owner
	for rows.Next() {
		if err := rows.Scan(
			&found.ID,
			&found.ManagerID,
			&found.Fullname,
			&found.CPF.Cpf,
			&found.RG.Rg,
			&found.Phone.Number,
			&found.Email.Address,
			&found.Occupation,
			&found.MaritalStatus,
			&found.Address.FullAddress,
			&found.Address.Street,
			&found.Address.Number,
			&found.Address.Complement,
			&found.Address.Neighborhood,
			&found.Address.City,
			&found.Address.State,
			&found.Address.ZipCode,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, httperr.NewNotFoundError(ctx, "owner not found or not exists")
			}

			return nil, httperr.NewInternalServerError(ctx, err.Error())
		}

		owners = append(owners, found)
	}

	return owners, nil
}
