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

func (r *InMemoryPropertyRepository) FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err) {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, jwt.ERR_TOKEN_INVALID_OR_EXPIRED)
	}

	found := []property.DTO{}

	for _, property := range r.properties {
		if property.UserID == uuid.MustParse(user_id) {
			found = append(found, *property.ToDTO())
		}
	}

	return found, nil
}
