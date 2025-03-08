package environment

import (
	"github.com/ogabrielrodrigues/relatorio_saaec/kind"
)

func LoadEnvironment() *kind.Environment {
	env := kind.Environment{}

	env.ReadEnvironmentFromStdin()

	return &env
}
