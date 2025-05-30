package owner

import (
	"context"

	"imobiliary/internal/entity/owner"
	"imobiliary/internal/response"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, dto owner.CreateDTO) (uuid.UUID, *response.Err) {
	new_owner, err := owner.New(dto.Fullname, dto.CPF, dto.RG, dto.Email, dto.Cellphone, dto.Occupation, dto.MaritalStatus, dto.Address.ToAddress(), uuid.Nil)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(ctx, *new_owner)
}
