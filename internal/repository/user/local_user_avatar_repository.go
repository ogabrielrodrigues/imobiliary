package repository

import (
	"context"
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
	user_folder := filepath.Join(r.path, id)
	if _, err := os.Stat(user_folder); os.IsNotExist(err) {
		if err := os.Mkdir(user_folder, os.ModePerm); err != nil {
			return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
	}

	avatar_path := filepath.Join(user_folder, r.default_filename)

	if _, err := os.Stat(avatar_path); os.IsNotExist(err) {
		dst, err := os.Create(avatar_path)
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

	dst, err := os.OpenFile(avatar_path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	_, err = io.Copy(dst, avatar)
	if err != nil {
		return response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return nil
}
