package handler

import (
	"fmt"
	"manager/internal/config"
)

type Handler struct {
	config *config.Config
}

func NewHandler(config *config.Config) *Handler {
	return &Handler{config: config}
}

func (h *Handler) Handle() {
	fmt.Println("Handler initialized")
}
