package handler

import (
	"manager/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct {
	service *service.Service
}

type addExpenseRequest struct {
	SMS string `json:"sms"`
}

func NewExpenseHandler(service *service.Service) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

func (e *ExpenseHandler) AddExpense(c echo.Context) error {
	var req addExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if req.SMS == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "sms is required"})
	}

	expense, err := e.service.Expense().CreateExpense(req.SMS)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, expense)
}
