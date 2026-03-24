// Package handler
package handler

import (
	"net/http"
	"strconv"
	"time"

	"manager/internal/service"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service *service.Service
}

type addTransactionRequest struct {
	SMS string `json:"sms"`
}

type UpdateTransactionRequest struct {
	ID  string `json:"ID"`
	SMS string `json:"sms"`
}

const transactionDateLayout = "2006-01-02"

func NewTransactionHandler(service *service.Service) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (e *TransactionHandler) AddTransaction(c echo.Context) error {
	var req addTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if req.SMS == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "sms is required"})
	}

	transaction, err := e.service.Transaction().CreateTransaction(req.SMS)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, transaction)
}

func (e *TransactionHandler) UpdateTransaction(c echo.Context) error {
	var req UpdateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if req.ID == "" || req.SMS == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id and sms are required"})
	}

	transaction, err := e.service.Transaction().UpdateTransaction(req.ID, req.SMS)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

		return c.JSON(http.StatusOK, transaction)
}

func (e *TransactionHandler) GetTransactions(c echo.Context) error {
	transactions, err := e.service.Transaction().GetTransactions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}

func (e *TransactionHandler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	err := e.service.Transaction().DeleteTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (e *TransactionHandler) GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	transaction, err := e.service.Transaction().GetTransactionByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transaction)
}


func (e *TransactionHandler) GetTransactionsByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	if category == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "category is required"})
	}

	transactions, err := e.service.Transaction().GetTransactionsByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, transactions)
}

func (e *TransactionHandler) GetTransactionsByMerchant(c echo.Context) error {
	merchant := c.QueryParam("merchant")
	if merchant == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "merchant is required"})
	}

	transactions, err := e.service.Transaction().GetTransactionsByMerchant(merchant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, transactions)
}

func (e *TransactionHandler) GetTransactionsByDate(c echo.Context) error {
	dateString := c.QueryParam("date")
	if dateString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "date is required"})
	}
	date, err := time.Parse(transactionDateLayout, dateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "date must be in YYYY-MM-DD format"})
	}

	transactions, err := e.service.Transaction().GetTransactionsByDate(date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, transactions)
}

func (e *TransactionHandler) GetTransactionsByMonth(c echo.Context) error {
	yearString := c.QueryParam("year")
	monthString := c.QueryParam("month")
	if yearString == "" || monthString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "year and month are required"})
	}
	year, err := strconv.Atoi(yearString)
	if err != nil || year < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "year must be a valid positive integer"})
	}
	month, err := strconv.Atoi(monthString)
	if err != nil || month < 1 || month > 12 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "month must be an integer between 1 and 12"})
	}

	transactions, err := e.service.Transaction().GetTransactionsByMonth(year, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, transactions)
}


func (e *TransactionHandler) GetTransactionsByDateRange(c echo.Context) error {
	startString := c.QueryParam("start")
	endString := c.QueryParam("end")
	if startString == "" || endString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "start and end are required"})
	}
	start, err := time.Parse(transactionDateLayout, startString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "start must be in YYYY-MM-DD format"})
	}
		end, err := time.Parse(transactionDateLayout, endString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "end must be in YYYY-MM-DD format"})
	}
	if end.Before(start) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "end must be after start"})
	}

	transactions, err := e.service.Transaction().GetTransactionsByDateRange(start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
		return c.JSON(http.StatusOK, transactions)
}