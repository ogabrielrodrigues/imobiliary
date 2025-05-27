package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/ogabrielrodrigues/imobiliary/internal/domain/types"
)

type Config struct {
	environment        types.Environment
	serverAddr         string
	corsOrigin         string
	jwtSecret          string
	postgresConnString string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env := os.Getenv("ENVIRONMENT")

	var environment types.Environment
	switch env {
	case "production":
		environment = types.Production
	case "staging":
		environment = types.Staging
	default:
		environment = types.Development
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		return nil, errors.New("SERVER_ADDR env variable not defined")
	}

	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		return nil, errors.New("CORS_ORIGIN env variable not defined")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET env variable not defined")
	}

	postgresConnString := os.Getenv("DATABASE_URL")
	if postgresConnString == "" {
		return nil, errors.New("DATABASE_URL env variable not defined")
	}

	return &Config{
		environment:        environment,
		serverAddr:         serverAddr,
		corsOrigin:         corsOrigin,
		jwtSecret:          jwtSecret,
		postgresConnString: postgresConnString,
	}, nil
}

func (c *Config) GetServerAddr() string {
	return c.serverAddr
}

func (c *Config) GetCorsOrigin() string {
	return c.corsOrigin
}

func (c *Config) GetJwtSecret() string {
	return c.jwtSecret
}

func (c *Config) GetPostgresConnString() string {
	return c.postgresConnString
}
