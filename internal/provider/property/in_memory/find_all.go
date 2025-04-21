package provider

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *MemPropertyRepository) FindAllByUserID(ctx context.Context) ([]property.DTO, *response.Err) {
	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	found := []property.DTO{}

	for _, property := range r.properties {
		if property.UserID == uuid.MustParse(user_id) {
			found = append(found, *property.ToDTO())
		}
	}

	return found, nil
}
