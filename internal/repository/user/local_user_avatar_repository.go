package repository

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type LocalUserAvatarRepository struct {
	path             string
	default_filename string
}

func NewLocalUserAvatarRepository(path string) *LocalUserAvatarRepository {
	return &LocalUserAvatarRepository{
		path:             filepath.Join(path),
		default_filename: "avatar.png",
	}
}

func (r *LocalUserAvatarRepository) GetAvatar(ctx context.Context, id string) (string, *response.Err) {
	avatar_path := filepath.Join(r.path, id, r.default_filename)

	if _, err := os.Stat(avatar_path); os.IsNotExist(err) {
		return "", response.NewErr(http.StatusNotFound, user.ERR_AVATAR_NOT_FOUND)
	}

	return avatar_path, nil
}

func (r *LocalUserAvatarRepository) SaveAvatar(ctx context.Context, id string, avatar multipart.File) *response.Err {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", r.path, id)); os.IsNotExist(err) {
		if err := os.Mkdir(fmt.Sprintf("%s/%s", r.path, id), os.ModePerm); err != nil {
			return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s/%s", r.path, id, r.default_filename)); os.IsNotExist(err) {
		dst, err := os.Create(fmt.Sprintf("%s/%s/%s", r.path, id, r.default_filename))
		if err != nil {
			return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		defer dst.Close()

		_, err = io.Copy(dst, avatar)
		if err != nil {
			return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return nil
	}

	dst, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", r.path, id, r.default_filename), os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	_, err = io.Copy(dst, avatar)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return nil
}
