package environment

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	types "github.com/ogabrielrodrigues/imobiliary/internal/types/config"
)

var Environment *types.Environment

func Load() *types.Environment {
	if err := godotenv.Load(); err != nil {
		logger.Error("error loading .env file")
		os.Exit(1)
	}

	Environment = &types.Environment{
		SERVER_PROTOCOL: os.Getenv("SERVER_PROTOCOL"),
		SERVER_ADDR:     os.Getenv("SERVER_ADDR"),
		DATABASE_HOST:   os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:   os.Getenv("DATABASE_PORT"),
		DATABASE_NAME:   os.Getenv("DATABASE_NAME"),
		DATABASE_USER:   os.Getenv("DATABASE_USER"),
		DATABASE_PWD:    os.Getenv("DATABASE_PWD"),
		SECRET_KEY:      os.Getenv("SECRET_KEY"),
		CORS_ORIGIN:     os.Getenv("CORS_ORIGIN"),
	}

	return Environment
}

func LoadFile(path string) *types.Environment {
	if err := godotenv.Load(path); err != nil {
		logger.Error("error loading .env file")
		os.Exit(1)
	}

	Environment = &types.Environment{
		SERVER_PROTOCOL: os.Getenv("SERVER_PROTOCOL"),
		SERVER_ADDR:     os.Getenv("SERVER_ADDR"),
		DATABASE_HOST:   os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:   os.Getenv("DATABASE_PORT"),
		DATABASE_NAME:   os.Getenv("DATABASE_NAME"),
		DATABASE_USER:   os.Getenv("DATABASE_USER"),
		DATABASE_PWD:    os.Getenv("DATABASE_PWD"),
		SECRET_KEY:      os.Getenv("SECRET_KEY"),
		CORS_ORIGIN:     os.Getenv("CORS_ORIGIN"),
	}

	return Environment
}
