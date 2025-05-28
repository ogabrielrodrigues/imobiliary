package maker

import (
	"imobiliary/internal/api/handler"
)

func MakeManagerHandler() (*handler.ManagerHandler, error) {
	// user_repo, ur_err := repository.NewPostgresUserRepository(pool)
	// if ur_err != nil {
	// 	return nil
	// }

	return handler.NewManagerHandler(), nil
}
