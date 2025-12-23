package maker

import (
	"imobiliary/internal/api/handler"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/mail"
	"imobiliary/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeManagerHandler(pool *pgxpool.Pool) *handler.ManagerHandler {
	managerRepository := postgres.NewPostgresManagerRepository(pool)

	mailClient, _ := mail.NewMailClient()
	managerMailRepository := mail.NewMailManagerRepository(mailClient)

	managerHandler := handler.NewManagerHandler(
		managerRepository,
		manager.NewManagerValidator(),
		&managerMailRepository,
	)

	return managerHandler
}
