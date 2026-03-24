package handler

import (
	"manager/internal/service"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	healthHandler      *HealthHandler
	transactionHandler *TransactionHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		healthHandler:      NewHealthHandler(),
		transactionHandler: NewTransactionHandler(service),
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	transactions := e.Group("/transactions")
	e.GET("/", h.healthHandler.CheckHealth)
	transactions.POST("", h.transactionHandler.AddTransaction)
	transactions.GET("", h.transactionHandler.GetTransactions)
	transactions.GET("/:id", h.transactionHandler.GetTransactionByID)
	transactions.PUT("/:id", h.transactionHandler.UpdateTransaction)
	transactions.DELETE("/:id", h.transactionHandler.DeleteTransaction)
	transactions.GET("/category", h.transactionHandler.GetTransactionsByCategory)
	transactions.GET("/merchant", h.transactionHandler.GetTransactionsByMerchant)
	transactions.GET("/date", h.transactionHandler.GetTransactionsByDate)
	transactions.GET("/month", h.transactionHandler.GetTransactionsByMonth)
	transactions.GET("/date-range", h.transactionHandler.GetTransactionsByDateRange)
}
