package maker

import (
	"imobiliary/config"
	"imobiliary/internal/api/handler"
	"imobiliary/internal/application/usecase"
	"imobiliary/internal/domain/owner"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeOwnerHandler(pool *pgxpool.Pool, config *config.Config) *handler.OwnerHandler {
	ownerRepository := postgres.NewPostgresOwnerRepository(pool)

	findByIDOwnerUseCase := usecase.NewFindByIDOwner(ownerRepository)
	createOwnerUseCase := usecase.NewCreateOwner(ownerRepository, owner.NewOwnerValidator())
	findAllOwnerUseCase := usecase.NewFindAllOwner(ownerRepository)

	ownerHandler := handler.NewOwnerHandler(
		findByIDOwnerUseCase,
		createOwnerUseCase,
		findAllOwnerUseCase,
	)

	return ownerHandler
}
