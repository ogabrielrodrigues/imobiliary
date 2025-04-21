package factory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	property_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/handler"
	property_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/service"
	property_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/property/postgres"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func NewPropertyHandlerFactory(pool *pgxpool.Pool) (*property_handler.Handler, *response.Err) {
	property_repo, err := property_repository.NewPostgresPropertyRepository(pool)
	if err != nil {
		return nil, err
	}

	return property_handler.NewHandler(
		property_service.NewService(property_repo),
	), nil
}
