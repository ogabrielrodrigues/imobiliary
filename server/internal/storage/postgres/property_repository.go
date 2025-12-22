package postgres

import (
	"context"
	"database/sql"
	"errors"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/property"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPropertyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresPropertyRepository(pool *pgxpool.Pool) *PostgresPropertyRepository {
	return &PostgresPropertyRepository{db: pool}
}

func (pr *PostgresPropertyRepository) FindByID(ctx context.Context, propertyID, managerID uuid.UUID) (*property.Property, *httperr.HttpError) {
	row := pr.db.QueryRow(ctx, `
		select
			pr.*,
			ad.full_address,
			ad.street,
			ad.number,
			ad.complement,
			ad.neighborhood,
			ad.city,
			ad.state,
			ad.zip_code
		from "property" pr
		join "address" ad on pr.address_id = ad.id
		left join "owner" ow on pr.owner_id = ow.id
		where pr.id = $1 and pr.manager_id = $2 and pr.active = true`,
		propertyID,
		managerID,
	)

	var found property.Property
	if err := row.Scan(
		&found.ID,
		&found.Address.ID,
		&found.ManagerID,
		&found.OwnerID,
		&found.Status,
		&found.Kind,
		&found.WaterID,
		&found.EnergyID,
		&found.Active,
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
			return nil, httperr.NewNotFoundError(ctx, "property not found or not exists")
		}

		return nil, httperr.NewInternalServerError(ctx, err.Error())
	}

	return &found, nil
}

func (pr *PostgresPropertyRepository) Create(ctx context.Context, property *property.Property, managerID uuid.UUID) *httperr.HttpError {
	tx, err := pr.db.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error creating property address")
	}

	row := tx.QueryRow(ctx, `
		insert into "address" (full_address, street, number, complement, neighborhood, city, state, zip_code)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`,
		property.Address.FullAddress,
		property.Address.Street,
		property.Address.Number,
		property.Address.Complement,
		property.Address.Neighborhood,
		property.Address.City,
		property.Address.State,
		property.Address.ZipCode,
	)

	var addressID string
	if err := row.Scan(&addressID); err != nil {
		tx.Rollback(ctx)

		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "property address already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	_, err = tx.Exec(ctx, `
		insert into "property" values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		property.ID,
		addressID,
		managerID,
		property.OwnerID,
		property.Status,
		property.Kind,
		property.WaterID,
		property.EnergyID,
		property.Active,
	)

	if err != nil {
		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "property already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error committing transaction")
	}

	return nil
}

func (pr *PostgresPropertyRepository) FindAll(ctx context.Context, managerID uuid.UUID) ([]property.Property, *httperr.HttpError) {
	rows, err := pr.db.Query(ctx, `
		select
			pr.*,
			ad.full_address,
			ad.street,
			ad.number,
			ad.complement,
			ad.neighborhood,
			ad.city,
			ad.state,
			ad.zip_code
		from "property" pr
		join "address" ad on pr.address_id = ad.id
		left join "owner" ow on pr.owner_id = ow.id
		where pr.manager_id = $1 and pr.active = true`,
		managerID,
	)

	if err != nil {
		if IsErrNoRows(err) {
			return []property.Property{}, nil
		}

		return nil, httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	var found property.Property
	var properties []property.Property
	for rows.Next() {
		if err := rows.Scan(
			&found.ID,
			&found.Address.ID,
			&found.ManagerID,
			&found.OwnerID,
			&found.Status,
			&found.Kind,
			&found.WaterID,
			&found.EnergyID,
			&found.Active,
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
				return nil, httperr.NewNotFoundError(ctx, "property not found or not exists")
			}

			return nil, httperr.NewInternalServerError(ctx, err.Error())
		}

		properties = append(properties, found)
	}

	return properties, nil
}
