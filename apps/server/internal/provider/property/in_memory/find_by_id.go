package provider

import (
	"context"
	"net/http"

	"imobiliary/internal/entity/property"
	jwt "imobiliary/internal/lib"
	"imobiliary/internal/middleware"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (r *InMemoryPropertyRepository) FindByID(ctx context.Context, id uuid.UUID) (*property.DTO, *response.Err) {
	_, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	for _, property := range r.properties {
		if property.ID == id {
			return property.ToDTO(), nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
}
