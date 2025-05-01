package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (pg *PostgresOwnerRepository) Create(ctx context.Context, owner owner.Owner) (uuid.UUID, *response.Err) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	address_query := `
		INSERT INTO "address" (street, "number", complement, neighborhood, city, state, zip_code, full_address, mini_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	row := tx.QueryRow(ctx, address_query,
		owner.Address.Street,
		owner.Address.Number,
		owner.Address.Complement,
		owner.Address.Neighborhood,
		owner.Address.City,
		owner.Address.State,
		owner.Address.ZipCode,
		owner.Address.FullAddress,
		owner.Address.MiniAddress,
	)

	var address_id string
	if err := row.Scan(&address_id); err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	owner_query := `
		INSERT INTO "owner" (id, fullname, cpf, rg, email, cellphone, occupation, marital_status, address_id, manager_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;
	`

	manager_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, lib.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	row = tx.QueryRow(ctx, owner_query,
		owner.ID,
		owner.Fullname,
		owner.CPF,
		owner.RG,
		owner.Email,
		owner.Cellphone,
		owner.Occupation,
		owner.MaritalStatus,
		address_id,
		manager_id,
	)

	var owner_id string
	if err := row.Scan(&owner_id); err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, err.Error())
	}

	return uuid.MustParse(owner_id), nil
}
