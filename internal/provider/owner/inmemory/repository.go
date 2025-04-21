package provider

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/owner"
)

type InMemoryOwnerRepository struct {
	owners []owner.Owner
}

func NewInMemoryOwnerRepository() *InMemoryOwnerRepository {
	return &InMemoryOwnerRepository{
		owners: []owner.Owner{},
	}
}
