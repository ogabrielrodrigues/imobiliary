package factory

import (
	property_handler "imobiliary/internal/entity/property/handler"
	property_service "imobiliary/internal/entity/property/service"
	property_repository "imobiliary/internal/provider/property/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPropertyHandlerFactory(pool *pgxpool.Pool) *property_handler.Handler {
	property_repo, err := property_repository.NewPostgresPropertyRepository(pool)
	if err != nil {
		return nil
	}

	return property_handler.NewHandler(property_service.NewService(property_repo))
}
