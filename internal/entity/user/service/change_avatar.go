package user

import (
	"context"
	"mime/multipart"

	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) ChangeAvatar(ctx context.Context, avatar_file multipart.File, mime string) *response.Err {
	avatar_url, err := s.storage.ChangeAvatar(ctx, avatar_file, mime)
	if err != nil {
		return err
	}

	return s.repo.ChangeAvatar(ctx, avatar_url)
}
