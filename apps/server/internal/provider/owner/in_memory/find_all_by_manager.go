package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/owner"
	jwt "imobiliary/internal/lib"
	"imobiliary/internal/middleware"
	"imobiliary/internal/response"
)

func (r *InMemoryOwnerRepository) FindAllByManagerID(ctx context.Context) ([]owner.DTO, *response.Err) {
	manager_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	var owners []owner.DTO
	for _, o := range r.owners {
		if o.ManagerID.String() == manager_id {
			owners = append(owners, *o.ToDTO())
		}
	}

	return owners, nil
}
