package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr         string
	CorsOrigin         string
	JwtSecret          string
	PostgresConnString string
	MailAddr           string
	MailHost           string
	MailPort           string
	MailUser           string
	MailPassword       string
	ResetPasswordUrl   string
}

var globalConfig Config

func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error reading environment variables")
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		return fmt.Errorf("SERVER_ADDR environment variable is not defined")
	}

	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		return fmt.Errorf("CORS_ORIGIN environment variable is not defined")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return fmt.Errorf("JWT_SECRET environment variable is not defined")
	}

	postgresConnString := os.Getenv("DATABASE_URL")
	if postgresConnString == "" {
		return fmt.Errorf("DATABASE_URL env variable not defined")
	}

	mailAddr := os.Getenv("MAIL_ADDR")
	if mailAddr == "" {
		return fmt.Errorf("MAIL_ADDR env variable not defined")
	}

	mailHost := os.Getenv("MAIL_HOST")
	if mailHost == "" {
		return fmt.Errorf("MAIL_HOST env variable not defined")
	}

	mailPort := os.Getenv("MAIL_PORT")
	if mailPort == "" {
		return fmt.Errorf("MAIL_PORT env variable not defined")
	}

	mailUser := os.Getenv("MAIL_USER")
	if mailUser == "" {
		return fmt.Errorf("MAIL_USER env variable not defined")
	}

	mailPassword := os.Getenv("MAIL_PASSWORD")
	if mailPassword == "" {
		return fmt.Errorf("MAIL_PASSWORD env variable not defined")
	}

	resetPasswordUrl := os.Getenv("RESET_PASSWORD_URL")
	if resetPasswordUrl == "" {
		return fmt.Errorf("RESET_PASSWORD_URL env variable not defined")
	}

	globalConfig.ServerAddr = serverAddr
	globalConfig.CorsOrigin = corsOrigin
	globalConfig.JwtSecret = jwtSecret
	globalConfig.PostgresConnString = postgresConnString
	globalConfig.MailAddr = mailAddr
	globalConfig.MailHost = mailHost
	globalConfig.MailPort = mailPort
	globalConfig.MailUser = mailUser
	globalConfig.MailPassword = mailPassword
	globalConfig.ResetPasswordUrl = resetPasswordUrl

	return nil
}

func GetConfig() *Config {
	return &globalConfig
}
