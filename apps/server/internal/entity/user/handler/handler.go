package user

import (
	user "imobiliary/internal/entity/user/service"
)

type Handler struct {
	service user.IService
}

func NewHandler(service user.IService) *Handler {
	return &Handler{service}
}
