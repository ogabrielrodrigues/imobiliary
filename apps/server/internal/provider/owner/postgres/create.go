package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/owner"
	jwt "imobiliary/internal/lib"
	"imobiliary/internal/middleware"
	"imobiliary/internal/response"
	"imobiliary/internal/store"

	"github.com/google/uuid"
)

func (pg *PostgresOwnerRepository) Create(ctx context.Context, dto owner.Owner) (uuid.UUID, *response.Err) {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	address_query := `
		INSERT INTO "address" (street, "number", complement, neighborhood, city, state, zip_code, full_address, mini_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	row := tx.QueryRow(ctx, address_query,
		dto.Address.Street,
		dto.Address.Number,
		dto.Address.Complement,
		dto.Address.Neighborhood,
		dto.Address.City,
		dto.Address.State,
		dto.Address.ZipCode,
		dto.Address.FullAddress,
		dto.Address.MiniAddress,
	)

	var address_id string
	if err := row.Scan(&address_id); err != nil {
		tx.Rollback(ctx)

		if store.IsUniqueConstraint(err) {
			return uuid.Nil, response.NewErr(http.StatusConflict, owner.ERR_OWNER_ALREADY_EXISTS)
		}

		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	owner_query := `
		INSERT INTO "owner" (id, fullname, cpf, rg, email, cellphone, occupation, marital_status, address_id, manager_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;
	`

	manager_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return uuid.Nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	row = tx.QueryRow(ctx, owner_query,
		dto.ID,
		dto.Fullname,
		dto.CPF,
		dto.RG,
		dto.Email,
		dto.Cellphone,
		dto.Occupation,
		dto.MaritalStatus,
		address_id,
		manager_id,
	)

	var owner_id string
	if err := row.Scan(&owner_id); err != nil {
		tx.Rollback(ctx)

		if store.IsUniqueConstraint(err) {
			return uuid.Nil, response.NewErr(http.StatusConflict, owner.ERR_OWNER_ALREADY_EXISTS)
		}

		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return uuid.Nil, response.NewErr(http.StatusInternalServerError, response.ERR_INTERNAL_SERVER_ERROR)
	}

	return uuid.MustParse(owner_id), nil
}
