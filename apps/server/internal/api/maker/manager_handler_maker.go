package maker

import (
	"imobiliary/config"
	"imobiliary/internal/api/handler"
	"imobiliary/internal/application/usecase"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeManagerHandler(pool *pgxpool.Pool, config *config.Config) (*handler.ManagerHandler, error) {
	managerRepository := postgres.NewPostgresManagerRepository(pool)

	findByIDManagerUseCase := usecase.NewFindByIDManager(managerRepository)
	createManagerUseCase := usecase.NewCreateManager(managerRepository, manager.NewManagerValidator())
	authenticateManagerUseCase := usecase.NewAuthenticateManager(managerRepository, config.GetJwtSecret())

	return handler.NewManagerHandler(findByIDManagerUseCase,
		createManagerUseCase, authenticateManagerUseCase), nil
}
