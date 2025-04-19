package user

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/lib"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) SaveAvatar(ctx context.Context, avatarFile multipart.File) (string, *response.Err) {
	avatar_url, err := s.storage.SaveAvatar(ctx, avatarFile)
	if err != nil {
		return "", err
	}

	user_id, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return "", response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	err = s.repo.ChangeAvatar(ctx, avatar_url)
	if err != nil {
		return "", err
	}

	found, err := s.FindByID(ctx, uuid.MustParse(user_id))
	if err != nil {
		return "", err
	}

	return lib.GenerateToken(found.ToUser())
}
