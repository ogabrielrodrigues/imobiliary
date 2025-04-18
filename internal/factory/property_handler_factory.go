package factory

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/property"
	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/repository/property"
)

func NewPropertyHandlerFactory() *property.Handler {
	property_repo := property_repository.NewMemPropertyRepository()

	return property.NewHandler(
		property.NewService(property_repo),
	)
}
