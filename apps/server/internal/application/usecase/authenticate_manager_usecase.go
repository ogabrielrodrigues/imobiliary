package usecase

import (
	"context"
	"imobiliary/internal/application/dto/request"
	"imobiliary/internal/application/jwt"
	"imobiliary/internal/domain/manager"
	"imobiliary/internal/domain/types"
)

type AuthenticateManager struct {
	repository manager.Repository
}

func NewAuthenticateManager(repository manager.Repository) *AuthenticateManager {
	return &AuthenticateManager{repository}
}

func (cm *AuthenticateManager) Execute(ctx context.Context, dto request.AuthDTO) (string, error) { // TODO: place error type
	email, err := types.NewEmail(dto.Email)
	if err != nil {
		return "", err
	}

	managerID, err := cm.repository.Authenticate(ctx, email, dto.Password)
	if err != nil {
		return "", err // TODO: place error type
	}

	token, err := jwt.GenerateToken(managerID)
	if err != nil {
		return "", err // TODO: place error type
	}

	return token, nil
}
