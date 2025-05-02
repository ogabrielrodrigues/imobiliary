package user

import (
	"net/http"

	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

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

func (h *Handler) ChangeAvatar(w http.ResponseWriter, r *http.Request) *response.Err {
	file, metadata, err := r.FormFile("avatar")
	if err != nil {
		return response.NewErr(http.StatusBadRequest, user.ERR_MUST_BE_PROVIDE_AVATAR)
	}
	defer file.Close()

	if metadata.Size > 3*1024*1024 {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_SIZE_INVALID)
	}

	file_mime := metadata.Header.Get("Content-Type")
	if !haveMimeType(file_mime) {
		return response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_FORMAT_INVALID)
	}

	if err := h.service.ChangeAvatar(r.Context(), file, file_mime); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
