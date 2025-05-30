package factory

import (
	owner_handler "imobiliary/internal/entity/owner/handler"
	owner_service "imobiliary/internal/entity/owner/service"
	owner_repository "imobiliary/internal/provider/owner/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewOwnerHandlerFactory(pool *pgxpool.Pool) *owner_handler.Handler {
	owner_repo, err := owner_repository.NewPostgresOwnerRepository(pool)
	if err != nil {
		return nil
	}

	return owner_handler.NewHandler(owner_service.NewService(owner_repo))
}
