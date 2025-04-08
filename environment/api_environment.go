package environment

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ogabrielrodrigues/imobiliary/internal/kind"
	"github.com/ogabrielrodrigues/imobiliary/internal/shared"
)

func LoadAPIEnvironment() *kind.APIEnvironment {
	env := kind.APIEnvironment{}

	if err := godotenv.Load(); err != nil {
		shared.Logln(shared.ColorRed, "✗ ERROR reading .env")
		os.Exit(1)
	}

	env.SERVER_ADDR = os.Getenv("SERVER_ADDR")
	env.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	env.DATABASE_PORT = os.Getenv("DATABASE_PORT")
	env.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	env.DATABASE_USER = os.Getenv("DATABASE_USER")
	env.DATABASE_PWD = os.Getenv("DATABASE_PWD")

	return &env
}
