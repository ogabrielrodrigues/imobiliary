package user

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (s *Service) ChangeAvatar(ctx context.Context, file multipart.File, metadata *multipart.FileHeader) *response.Err {
	defer file.Close()

	if metadata.Size > 3*1024*1024 {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_SIZE_INVALID)
	}

	file_mime := metadata.Header.Get("Content-Type")
	if !haveMimeType(file_mime) {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_FORMAT_INVALID)
	}

	avatar_url, err := s.storage.ChangeAvatar(ctx, file, mime)
	if err != nil {
		return err
	}

	return s.repo.ChangeAvatar(ctx, avatar_url)
}
