package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	jwt "github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (r *InMemoryOwnerRepository) AssignOwnerToProperty(ctx context.Context, owner_id uuid.UUID, property_id uuid.UUID) *response.Err {
	_, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	_, err := r.FindByID(ctx, owner_id)
	if err != nil {
		return err
	}

	for _, found := range r.properties {
		if found.ID == property_id {
			found.OwnerID = owner_id
			return nil
		}
	}

	return response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
}
