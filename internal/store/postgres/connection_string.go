package store

import (
	"fmt"

	"github.com/ogabrielrodrigues/imobiliary/internal/types"
)

func PostgresConnectionString(env types.Environment) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.DATABASE_USER,
		env.DATABASE_PWD,
		env.DATABASE_HOST,
		env.DATABASE_PORT,
		env.DATABASE_NAME,
	)
}
