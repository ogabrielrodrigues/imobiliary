package environment

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	types "github.com/ogabrielrodrigues/imobiliary/internal/types/config"
)

const (
	ERR_READING_ENV_FILE = "error loading .env file"
)

var Environment *types.Environment

func Load() *types.Environment {
	if err := godotenv.Load(); err != nil {
		logger.Log(ERR_READING_ENV_FILE)
		os.Exit(1)
	}

	Environment = &types.Environment{
		SERVER_PROTOCOL:  os.Getenv("SERVER_PROTOCOL"),
		SERVER_ADDR:      os.Getenv("SERVER_ADDR"),
		SECRET_KEY:       os.Getenv("SECRET_KEY"),
		CORS_ORIGIN:      os.Getenv("CORS_ORIGIN"),
		S3_PUBLIC_URL:    os.Getenv("S3_PUBLIC_URL"),
		S3_AVATAR_BUCKET: os.Getenv("S3_AVATAR_BUCKET"),
		S3_TOKEN:         os.Getenv("S3_TOKEN"),
		S3_ACCESS_KEY:    os.Getenv("S3_ACCESS_KEY"),
		S3_SECRET_KEY:    os.Getenv("S3_SECRET_KEY"),
		S3_ACCOUNT_ID:    os.Getenv("S3_ACCOUNT_ID"),
		DATABASE_HOST:    os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:    os.Getenv("DATABASE_PORT"),
		DATABASE_NAME:    os.Getenv("DATABASE_NAME"),
		DATABASE_USER:    os.Getenv("DATABASE_USER"),
		DATABASE_PWD:     os.Getenv("DATABASE_PWD"),
	}

	return Environment
}

func LoadFile(path string) *types.Environment {
	if err := godotenv.Load(path); err != nil {
		logger.Log(ERR_READING_ENV_FILE)
		os.Exit(1)
	}

	Environment = &types.Environment{
		SERVER_PROTOCOL:  os.Getenv("SERVER_PROTOCOL"),
		SERVER_ADDR:      os.Getenv("SERVER_ADDR"),
		SECRET_KEY:       os.Getenv("SECRET_KEY"),
		CORS_ORIGIN:      os.Getenv("CORS_ORIGIN"),
		S3_PUBLIC_URL:    os.Getenv("S3_PUBLIC_URL"),
		S3_AVATAR_BUCKET: os.Getenv("S3_AVATAR_BUCKET"),
		S3_TOKEN:         os.Getenv("S3_TOKEN"),
		S3_ACCESS_KEY:    os.Getenv("S3_ACCESS_KEY"),
		S3_SECRET_KEY:    os.Getenv("S3_SECRET_KEY"),
		S3_ACCOUNT_ID:    os.Getenv("S3_ACCOUNT_ID"),
		DATABASE_HOST:    os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:    os.Getenv("DATABASE_PORT"),
		DATABASE_NAME:    os.Getenv("DATABASE_NAME"),
		DATABASE_USER:    os.Getenv("DATABASE_USER"),
		DATABASE_PWD:     os.Getenv("DATABASE_PWD"),
	}

	return Environment
}
