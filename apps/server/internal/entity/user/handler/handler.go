package user

import (
	user "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
)

type Handler struct {
	service user.IService
}

func NewHandler(service user.IService) *Handler {
	return &Handler{service}
}
