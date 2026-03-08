package handler

import (
	"manager/internal/service"
	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct {
	service *service.Service
}

func NewExpenseHandler(service *service.Service) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

func (e *ExpenseHandler) AddExpense(c echo.Context){
	var req
}
