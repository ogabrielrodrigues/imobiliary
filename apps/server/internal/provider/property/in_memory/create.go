package provider

import (
	"context"

	"imobiliary/internal/entity/property"
	"imobiliary/internal/response"
)

func (r *InMemoryPropertyRepository) Create(ctx context.Context, property *property.Property) *response.Err {
	r.properties = append(r.properties, property)
	return nil
}
