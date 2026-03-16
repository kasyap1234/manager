// Package handler
package handler

import (
	"net/http"
	"strconv"
	"time"

	"manager/internal/service"

	"github.com/labstack/echo/v4"
)

type ExpenseHandler struct {
	service *service.Service
}

type addExpenseRequest struct {
	SMS string `json:"sms"`
}

type UpdateExpenseRequest struct {
	ID  string `json:"ID"`
	SMS string `json:"sms"`
}

const expenseDateLayout = "2006-01-02"

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

func (e *ExpenseHandler) UpdateExpense(c echo.Context) error {
	var req UpdateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if req.ID == "" || req.SMS == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id and sms are required"})
	}

	expense, err := e.service.Expense().UpdateExpense(req.ID, req.SMS)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}

func (e *ExpenseHandler) GetExpenses(c echo.Context) error {
	expenses, err := e.service.Expense().GetExpenses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) DeleteExpense(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	err := e.service.Expense().DeleteExpense(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (e *ExpenseHandler) GetExpenseByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	expense, err := e.service.Expense().GetExpenseByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, expense)
}


func (e *ExpenseHandler) GetExpensesByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	if category == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "category is required"})
	}

	expenses, err := e.service.Expense().GetExpensesByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) GetExpensesByMerchant(c echo.Context) error {
	merchant := c.QueryParam("merchant")
	if merchant == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "merchant is required"})
	}

	expenses, err := e.service.Expense().GetExpensesByMerchant(merchant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) GetExpensesByDate(c echo.Context) error {
	dateString := c.QueryParam("date")
	if dateString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "date is required"})
	}
	date, err := time.Parse(expenseDateLayout, dateString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "date must be in YYYY-MM-DD format"})
	}

	expenses, err := e.service.Expense().GetExpensesByDate(date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) GetExpensesByMonth(c echo.Context) error {
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

	expenses, err := e.service.Expense().GetExpensesByMonth(year, month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}


func (e *ExpenseHandler) GetExpensesByDateRange(c echo.Context) error {
	startString := c.QueryParam("start")
	endString := c.QueryParam("end")
	if startString == "" || endString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "start and end are required"})
	}
	start, err := time.Parse(expenseDateLayout, startString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "start must be in YYYY-MM-DD format"})
	}
	end, err := time.Parse(expenseDateLayout, endString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "end must be in YYYY-MM-DD format"})
	}
	if end.Before(start) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "end must be after start"})
	}

	expenses, err := e.service.Expense().GetExpensesByDateRange(start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, expenses)
}