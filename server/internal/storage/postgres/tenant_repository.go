package postgres

import (
	"context"
	"database/sql"
	"errors"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/domain/tenant"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresTenantRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTenantRepository(pool *pgxpool.Pool) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: pool}
}

func (tr *PostgresTenantRepository) FindByID(ctx context.Context, tenantID uuid.UUID, managerID uuid.UUID) (*tenant.Tenant, *httperr.HttpError) {
	row := tr.db.QueryRow(ctx, `
		select
			te.*,
			ad.full_address,
			ad.street,
			ad.number,
			ad.complement,
			ad.neighborhood,
			ad.city,
			ad.state,
			ad.zip_code
		from "tenant" te
		join "address" ad on te.address_id = ad.id
		where te.id = $1 and te.manager_id = $2 and te.active = true`,
		tenantID,
		managerID,
	)

	var found tenant.Tenant
	if err := row.Scan(
		&found.ID,
		&found.ManagerID,
		&found.Address.ID,
		&found.Fullname,
		&found.CPF.Cpf,
		&found.RG.Rg,
		&found.Phone.Number,
		&found.Occupation,
		&found.MaritalStatus,
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
			return nil, httperr.NewNotFoundError(ctx, "tenant not found or not exists")
		}

		return nil, httperr.NewInternalServerError(ctx, err.Error())
	}

	return &found, nil
}

func (tr *PostgresTenantRepository) Create(ctx context.Context, tenant *tenant.Tenant, managerID uuid.UUID) *httperr.HttpError {
	tx, err := tr.db.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error creating tenant address")
	}

	row := tx.QueryRow(ctx, `
		insert into "address" (full_address, street, number, complement, neighborhood, city, state, zip_code)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`,
		tenant.Address.FullAddress,
		tenant.Address.Street,
		tenant.Address.Number,
		tenant.Address.Complement,
		tenant.Address.Neighborhood,
		tenant.Address.City,
		tenant.Address.State,
		tenant.Address.ZipCode,
	)

	var addressID string
	if err := row.Scan(&addressID); err != nil {
		tx.Rollback(ctx)

		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "tenant address already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	_, err = tx.Exec(ctx, `
		insert into "tenant" values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		tenant.ID,
		managerID,
		addressID,
		tenant.Fullname,
		tenant.CPF.Value(),
		tenant.RG.Value(),
		tenant.Phone.Value(),
		tenant.Occupation,
		tenant.MaritalStatus,
		tenant.Active,
	)

	if err != nil {
		if IsUniqueConstraint(err) {
			return httperr.NewAlreadyExistsError(ctx, "tenant already exists")
		}

		return httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return httperr.NewInternalServerError(ctx, "error committing transaction")
	}

	return nil
}

func (tr *PostgresTenantRepository) FindAll(ctx context.Context, managerID uuid.UUID) ([]tenant.Tenant, *httperr.HttpError) {
	rows, err := tr.db.Query(ctx, `
		select
			te.*,
			ad.full_address,
			ad.street,
			ad.number,
			ad.complement,
			ad.neighborhood,
			ad.city,
			ad.state,
			ad.zip_code
		from "tenant" te
		join "address" ad on te.address_id = ad.id
		where te.manager_id = $1 and te.active = true`,
		managerID,
	)

	if err != nil {
		if IsErrNoRows(err) {
			return []tenant.Tenant{}, nil
		}

		return nil, httperr.NewInternalServerError(ctx, httperr.InternalServerError)
	}

	var found tenant.Tenant
	var tenants []tenant.Tenant
	for rows.Next() {
		if err := rows.Scan(
			&found.ID,
			&found.ManagerID,
			&found.Address.ID,
			&found.Fullname,
			&found.CPF.Cpf,
			&found.RG.Rg,
			&found.Phone.Number,
			&found.Occupation,
			&found.MaritalStatus,
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
				return nil, httperr.NewNotFoundError(ctx, "tenant not found or not exists")
			}

			return nil, httperr.NewInternalServerError(ctx, err.Error())
		}

		tenants = append(tenants, found)
	}

	return tenants, nil
}
