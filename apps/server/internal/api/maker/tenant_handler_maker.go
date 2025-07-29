package maker

import (
	"imobiliary/config"
	"imobiliary/internal/api/handler"
	"imobiliary/internal/application/usecase"
	"imobiliary/internal/domain/tenant"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeTenantHandler(pool *pgxpool.Pool, config *config.Config) *handler.TenantHandler {
	tenantRepository := postgres.NewPostgresTenantRepository(pool)

	findByIDTenantUseCase := usecase.NewFindByIDTenant(tenantRepository)
	createTenantUseCase := usecase.NewCreateTenant(tenantRepository, tenant.NewTenantValidator())
	findAllTenantUseCase := usecase.NewFindAllTenant(tenantRepository)

	tenantHandler := handler.NewTenantHandler(
		findByIDTenantUseCase,
		createTenantUseCase,
		findAllTenantUseCase,
	)

	return tenantHandler
}
