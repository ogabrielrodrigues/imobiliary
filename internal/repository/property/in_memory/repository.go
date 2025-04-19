package repository

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
)

type MemPropertyRepository struct {
	properties []*property.Property
}

func NewMemPropertyRepository() *MemPropertyRepository {
	return &MemPropertyRepository{}
}
