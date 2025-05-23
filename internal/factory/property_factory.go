package factory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	property_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/handler"
	property_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/service"
	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/property/postgres"
)

func NewPropertyHandlerFactory(pool *pgxpool.Pool) *property_handler.Handler {
	property_repo, err := property_repository.NewPostgresPropertyRepository(pool)
	if err != nil {
		return nil
	}

	return property_handler.NewHandler(property_service.NewService(property_repo))
}
