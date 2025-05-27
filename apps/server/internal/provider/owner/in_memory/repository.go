package provider

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
)

type InMemoryOwnerRepository struct {
	owners     []owner.Owner
	properties []*property.Property
}

func NewInMemoryOwnerRepository(properties []*property.Property) *InMemoryOwnerRepository {
	return &InMemoryOwnerRepository{
		owners:     []owner.Owner{},
		properties: properties,
	}
}
