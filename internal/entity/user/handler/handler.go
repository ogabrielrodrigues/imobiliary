package user

import (
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/plan"
	user "github.com/ogabrielrodrigues/imobiliary/internal/entity/user/service"
)

type Handler struct {
	service      user.IService
	plan_service plan.IService
}

func NewHandler(service user.IService, plan_service plan.IService) *Handler {
	return &Handler{
		service:      service,
		plan_service: plan_service,
	}
}
