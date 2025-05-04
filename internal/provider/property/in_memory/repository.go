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
