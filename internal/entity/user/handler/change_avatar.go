package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
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

func (h *Handler) ChangeAvatar(w http.ResponseWriter, r *http.Request) {
	file, metadata, err := r.FormFile("avatar")
	if err != nil {
		err := response.NewErr(http.StatusBadRequest, user.ERR_MUST_BE_PROVIDE_AVATAR)
		response.End(w, err.Code, err)
		return
	}
	defer file.Close()

	if metadata.Size > 3*1024*1024 {
		err := response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_SIZE_INVALID)
		response.End(w, err.Code, err)
		return
	}

	file_mime := metadata.Header.Get("Content-Type")
	if !haveMimeType(file_mime) {
		err := response.NewErr(http.StatusBadRequest, user.ERR_AVATAR_FORMAT_INVALID)
		response.End(w, err.Code, err)
		return
	}

	r_err := h.service.ChangeAvatar(r.Context(), file, file_mime)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUserPlan(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value(middleware.UserIDKey).(string)

	uid, err := uuid.Parse(user_id)
	if err != nil {
		r_err := response.NewErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		response.End(w, r_err.Code, r_err)
		return
	}

	plan, r_err := h.plan_service.GetUserPlan(context.Background(), uid)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	response.End(w, http.StatusOK, plan)
}
