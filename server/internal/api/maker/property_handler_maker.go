package maker

import (
	"imobiliary/internal/api/handler"
	"imobiliary/internal/application/usecase"
	"imobiliary/internal/domain/property"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakePropertyHandler(pool *pgxpool.Pool) *handler.PropertyHandler {
	propertyRepository := postgres.NewPostgresPropertyRepository(pool)

	findByIDPropertyUseCase := usecase.NewFindByIDProperty(propertyRepository)
	createPropertyUseCase := usecase.NewCreateProperty(propertyRepository, property.NewPropertyValidator())
	findAllPropertyUseCase := usecase.NewFindAllProperty(propertyRepository)

	propertyHandler := handler.NewPropertyHandler(
		findByIDPropertyUseCase,
		createPropertyUseCase,
		findAllPropertyUseCase,
	)

	return propertyHandler
}
