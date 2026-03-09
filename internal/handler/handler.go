package handler

import (
	"manager/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {

	expenseHandler *ExpenseHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{

		expenseHandler: NewExpenseHandler(service),
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.POST("/expenses", h.expenseHandler.AddExpense)
	e.GET("/expenses", h.expenseHandler.GetExpenses)
	e.GET("/expenses/:id", h.expenseHandler.GetExpenseByID)
	e.PUT("/expenses", h.expenseHandler.UpdateExpense)
	e.DELETE("/expenses/:id", h.expenseHandler.DeleteExpense)
}
