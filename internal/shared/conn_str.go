package shared

import (
	"fmt"

	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func ConnStr(env *kind.APIEnvironment) string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s",
		env.DATABASE_USER,
		env.DATABASE_PWD,
		env.DATABASE_HOST,
		env.DATABASE_NAME,
	)
}
