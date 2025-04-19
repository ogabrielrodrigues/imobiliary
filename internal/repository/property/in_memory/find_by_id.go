package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *MemPropertyRepository) FindByID(ctx context.Context, id uuid.UUID) (*property.Property, *response.Err) {
	for _, property := range r.properties {
		if property.ID == id {
			return property, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
}
