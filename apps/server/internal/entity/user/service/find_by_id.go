package user

import (
	"context"

	"imobiliary/internal/entity/user"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*user.DTO, *response.Err) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToDTO(), nil
}
