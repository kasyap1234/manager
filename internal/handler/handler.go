package handler

import (
	"manager/internal/service"
)

type Handler struct {
	expenseHandler *ExpenseHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		expenseHandler: NewExpenseHandler(service),
	}
}
