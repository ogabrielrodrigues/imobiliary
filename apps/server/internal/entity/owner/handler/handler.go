package owner

import owner "imobiliary/internal/entity/owner/service"

type Handler struct {
	service owner.IService
}

func NewHandler(service owner.IService) *Handler {
	return &Handler{service}
}
