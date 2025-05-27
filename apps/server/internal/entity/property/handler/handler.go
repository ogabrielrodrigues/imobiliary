package property

import (
	property_service "github.com/ogabrielrodrigues/imobiliary/internal/entity/property/service"
)

type Handler struct {
	service property_service.IService
}

func NewHandler(service property_service.IService) *Handler {
	return &Handler{service}
}
