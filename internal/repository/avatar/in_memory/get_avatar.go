package repository

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r *InMemoryAvatarRepository) GetAvatar(ctx context.Context, id string) (string, *response.Err) {
	avatar_path := filepath.Join(r.path, id, r.default_filename)

	if _, err := os.Stat(avatar_path); os.IsNotExist(err) {
		return "", response.NewErr(http.StatusNotFound, user.ERR_AVATAR_NOT_FOUND)
	}

	return avatar_path, nil
}
