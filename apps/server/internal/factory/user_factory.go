package factory

import (
	"imobiliary/config/environment"
	user_handler "imobiliary/internal/entity/user/handler"
	user_service "imobiliary/internal/entity/user/service"
	avatar_repository "imobiliary/internal/provider/avatar/cloudflare"
	user_repository "imobiliary/internal/provider/user/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewUserHandlerFactory(pool *pgxpool.Pool) *user_handler.Handler {
	env := environment.Environment

	user_repo, ur_err := user_repository.NewPostgresUserRepository(pool)
	if ur_err != nil {
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
		return nil
	}

	return user_handler.NewHandler(user_service.NewService(user_repo, avatar_repo))
}
