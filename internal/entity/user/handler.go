package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/config/environment"
	"github.com/ogabrielrodrigues/imobiliary/config/logger"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) FindBy(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ctx := context.Background()

	email := params.Get("email")
	id := params.Get("id")

	if email != "" && id != "" {
		err := response.NewErr(http.StatusBadRequest, ERR_ONLY_ONE_MUST_PARAMETER_MUST_BE_PROVIDED)
		response.End(w, err.Code, err)
		return
	}

	if id == "" {
		if match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !match {
			err := response.NewErr(http.StatusBadRequest, ERR_EMAIL_INVALID)
			response.End(w, err.Code, err)
			return
		}

		user, err := h.service.FindByEmail(ctx, email)
		if err != nil {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, http.StatusOK, user)
		return
	}

	if email == "" {
		uid, u_err := uuid.Parse(id)
		if u_err != nil {
			err := response.NewErr(http.StatusBadRequest, ERR_UUID_INVALID)
			response.End(w, err.Code, err)
			return
		}

		user, err := h.service.FindByID(ctx, uid)
		if err != nil {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, http.StatusOK, user)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		err := response.NewErr(http.StatusBadRequest, ERR_INVALID_USER_REQUEST_BODY)
		response.End(w, err.Code, err)
		return
	}

	id, err := h.service.Create(ctx, &dto)
	if err != nil {
		response.End(w, err.Code, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/users/%s", id))
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	h.service.Update(nil, &UpdateDTO{})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	h.service.Delete(nil, uuid.MustParse(r.PathValue("id")))
}

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var dto AuthDTO
	ctx := context.Background()

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		err := response.NewErr(http.StatusBadRequest, ERR_INVALID_USER_REQUEST_BODY)
		response.End(w, err.Code, err)
	}

	token, err := h.service.Authenticate(ctx, dto.Email, dto.Password)
	if err != nil {
		if err.Message == ERR_USER_NOT_FOUND_OR_NOT_EXISTS {
			response.End(w, err.Code, err)
			return
		}

		response.End(w, err.Code, err)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Header.Get("Authorization") == "" {
		err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		response.End(w, err.Code, err)
		return
	}

	authorization, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.Environment.SECRET_KEY), nil
	})
	if err != nil {
		err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		response.End(w, err.Code, err)
		return
	}

	user_id, err := token.Claims.GetSubject()
	if err != nil {
		err := response.NewErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		response.End(w, err.Code, err)
		return
	}

	file, _, err := r.FormFile("avatar")
	if err != nil {
		err := response.NewErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		response.End(w, err.Code, err)
		return
	}
	defer file.Close()

	r_err := h.service.SaveAvatar(ctx, uuid.MustParse(user_id), file)
	if r_err != nil {
		logger.Error("Error saving avatar", "error", r_err)
		response.End(w, r_err.Code, r_err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := r.PathValue("user_id")

	user_id, err := uuid.Parse(id)
	if err != nil {
		err := response.NewErr(http.StatusBadRequest, ERR_UUID_INVALID)
		response.End(w, err.Code, err)
		return
	}

	avatar_path, r_err := h.service.GetAvatar(ctx, user_id)
	if r_err != nil {
		response.End(w, r_err.Code, r_err)
		return
	}

	http.ServeFile(w, r, avatar_path)
}
