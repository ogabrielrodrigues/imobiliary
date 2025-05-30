package usecase

import (
	"context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/httperr"
	"imobiliary/internal/application/jwt"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"
)

type AuthenticateManager struct {
	repository manager.Repository
	jwtSecret  string
}

func NewAuthenticateManager(repository manager.Repository, jwtSecret string) *AuthenticateManager {
	return &AuthenticateManager{repository, jwtSecret}
}

func (am *AuthenticateManager) Execute(ctx context.Context, dto request.AuthDTO) (string, *httperr.HttpError) { // TODO: place error type
	email, err := types.NewEmail(dto.Email)
	if err != nil {
		return "", httperr.NewUnprocessableEntityError(ctx, err.Error())
	}

	managerID, err := am.repository.Authenticate(ctx, email, dto.Password)
	if err != nil {
		return "", err.(*httperr.HttpError)
	}

	token, err := jwt.GenerateToken(managerID, am.jwtSecret)
	if err != nil {
		return "", err.(*httperr.HttpError)
	}

	return token, nil
}
