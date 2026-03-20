package handler

import (
	"manager/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	healthHandler  *HealthHandler
	expenseHandler *ExpenseHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		
		expenseHandler: NewExpenseHandler(service),
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	expenses:=e.Group("/expenses")
	e.GET("/",h.healthHandler.CheckHealth)
	expenses.POST("/expenses", h.expenseHandler.AddExpense)
	expenses.GET("", h.expenseHandler.GetExpenses)
	expenses.GET("/:id", h.expenseHandler.GetExpenseByID)
	expenses.PUT("/:id", h.expenseHandler.UpdateExpense)
	expenses.DELETE("/:id", h.expenseHandler.DeleteExpense)
	expenses.GET("/category", h.expenseHandler.GetExpensesByCategory)
	expenses.GET("/merchant", h.expenseHandler.GetExpensesByMerchant)
	expenses.GET("/date", h.expenseHandler.GetExpensesByDate)
	expenses.GET("/month", h.expenseHandler.GetExpensesByMonth)
	expenses.GET("/date-range", h.expenseHandler.GetExpensesByDateRange)
}
