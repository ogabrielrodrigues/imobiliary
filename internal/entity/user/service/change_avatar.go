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

	mimetype := metadata.Header.Get("Content-Type")
	if !haveMimeType(mimetype) {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_FORMAT_INVALID)
	}

	avatar_url, err := s.storage.ChangeAvatar(ctx, file, mimetype)
	if err != nil {
		return err
	}

	return s.repo.ChangeAvatar(ctx, avatar_url)
}

func haveMimeType(mime_type string) bool {
	mime_types := map[string]struct{}{
		"image/jpeg": {},
		"image/png":  {},
		"image/jpg":  {},
		"image/webp": {},
	}

	_, ok := mime_types[mime_type]
	return ok
}
