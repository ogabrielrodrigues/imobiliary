package provider

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *MemPropertyRepository) Create(ctx context.Context, property *property.Property) (*property.Property, *response.Err) {
	r.properties = append(r.properties, property)
	return property, nil
}
