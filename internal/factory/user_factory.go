package factory

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	plan_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/plan/service"
	user_handler "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/handler"
	user_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
	avatar_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/avatar/cloudflare"
	plan_repository "github.com/ogabrielrodrigues/imobiliary/internal/provider/plan/postgres"
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

	plan_repo, pr_err := plan_repository.NewPostgresPlanRepository(pool)
	if pr_err != nil {
		logger.Panicf("err: %s", pr_err)
		return nil
	}

	return user_handler.NewHandler(
		user_service.NewService(user_repo, avatar_repo),
		plan_service.NewService(plan_repo),
	)
}
