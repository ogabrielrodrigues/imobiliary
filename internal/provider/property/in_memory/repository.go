package provider

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
)

type InMemoryPropertyRepository struct {
	properties []*property.Property
}

func NewInMemoryPropertyRepository() *InMemoryPropertyRepository {
	return &InMemoryPropertyRepository{}
}

func (r *InMemoryPropertyRepository) GetProperties() []*property.Property {
	return r.properties
}
