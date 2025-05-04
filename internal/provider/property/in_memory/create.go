package provider

import (
	"context"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (r *InMemoryPropertyRepository) Create(ctx context.Context, property *property.Property) *response.Err {
	r.properties = append(r.properties, property)
	return nil
}
