package user

import (
	"context"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, dto *user.CreateDTO) (uuid.UUID, *response.Err) {
	usr, err := user.New(dto.CreciID, dto.Fullname, dto.Cellphone, dto.Email, dto.Password)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(ctx, usr)
}
