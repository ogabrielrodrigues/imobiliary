package store

import (
	"fmt"

	types "github.com/ogabrielrodrigues/imobiliary/internal/types/config"
)

func ConnectionString(env types.Environment) string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s",
		env.DATABASE_USER,
		env.DATABASE_PWD,
		env.DATABASE_HOST,
		env.DATABASE_NAME,
	)
}
