package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type MemPropertyRepository struct {
	properties []*property.Property
}

func NewMemPropertyRepository() *MemPropertyRepository {
	return &MemPropertyRepository{}
}

func (r *MemPropertyRepository) FindAllByUserID(ctx context.Context, user_id uuid.UUID) ([]property.DTO, *response.Err) {
	found := []property.DTO{}

	for _, property := range r.properties {
		if property.UserID == user_id {
			found = append(found, *property.ToDTO())
		}
	}

	return found, nil
}

func (r *MemPropertyRepository) FindByID(ctx context.Context, id uuid.UUID) (*property.Property, *response.Err) {
	for _, property := range r.properties {
		if property.ID == id {
			return property, nil
		}
	}

	return nil, response.NewErr(http.StatusNotFound, property.ERR_PROPERTY_NOT_FOUND_OR_NOT_EXISTS)
}

func (r *MemPropertyRepository) Create(ctx context.Context, property *property.Property) (*property.Property, *response.Err) {
	r.properties = append(r.properties, property)
	return property, nil
}
