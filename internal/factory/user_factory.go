package factory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	user_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/handler"
	user_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
	avatar_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/avatar/cloudflare"
	user_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/user/postgres"
)

func NewUserHandlerFactory(pool *pgxpool.Pool) *user_handler.Handler {
	env := environment.Environment

	user_repo, ur_err := user_repository.NewPostgresUserRepository(pool)
	if ur_err != nil {
		logger.Panicf("err: %s", ur_err)
		return nil
	}

	avatar_repo, ar_err := avatar_repository.NewCloudflareR2AvatarRepository(
		env.S3_PUBLIC_URL,
		env.S3_AVATAR_BUCKET,
		env.S3_ACCESS_KEY,
		env.S3_SECRET_KEY,
		env.S3_ACCOUNT_ID,
	)
	if ar_err != nil {
		logger.Panicf("err: %s", ar_err)
		return nil
	}

	return user_handler.NewHandler(user_service.NewService(user_repo, avatar_repo))
}
