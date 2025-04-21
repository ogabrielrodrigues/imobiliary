package factory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	owner_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/owner/handler"
	owner_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/owner/service"
	owner_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/owner/postgres"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func NewOwnerHandlerFactory(pool *pgxpool.Pool) (*owner_handler.Handler, *response.Err) {
	owner_repo, err := owner_repository.NewPostgresOwnerRepository(pool)
	if err != nil {
		return nil, err
	}

	return owner_handler.NewHandler(
		owner_service.NewService(owner_repo),
	), nil
}
