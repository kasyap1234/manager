package handler

import (
	"fmt"
	"manager/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) Handle() {
	fmt.Println("Handler initialized")
}
