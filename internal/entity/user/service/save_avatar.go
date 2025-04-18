package user

import (
	"context"
	"mime/multipart"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) SaveAvatar(ctx context.Context, avatarFile multipart.File) *response.Err {
	avatar_url, err := s.storage.SaveAvatar(ctx, avatarFile)
	if err != nil {
		return err
	}

	return s.repo.ChangeAvatar(ctx, avatar_url)
}
