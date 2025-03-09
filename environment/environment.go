package environment

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
)

func LoadEnvironment() *kind.Environment {
	env := kind.Environment{}

	env.ReadEnvironment()

	return &env
}
