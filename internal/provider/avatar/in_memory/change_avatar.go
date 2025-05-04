package provider

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ogabrielrodrigues/imobiliary/internal/response"
)

func (r *InMemoryAvatarRepository) ChangeAvatar(ctx context.Context, avatar multipart.File, mime string) (string, *response.Err) {
	user_folder := filepath.Join(r.path, ctx.Value("user_id").(string))
	if _, err := os.Stat(user_folder); os.IsNotExist(err) {
		if err := os.Mkdir(user_folder, os.ModePerm); err != nil {
			return "", response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
	}

	avatar_path := filepath.Join(user_folder, r.default_filename)

	if _, err := os.Stat(avatar_path); os.IsNotExist(err) {
		dst, err := os.Create(avatar_path)
		if err != nil {
			return "", response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		defer dst.Close()

		_, err = io.Copy(dst, avatar)
		if err != nil {
			return "", response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		return avatar_path, nil
	}

	dst, err := os.OpenFile(avatar_path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	_, err = io.Copy(dst, avatar)
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return avatar_path, nil
}
