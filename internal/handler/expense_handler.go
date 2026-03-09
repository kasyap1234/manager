// Package handler
package handler

import (
	"net/http"

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
	var req *UpdateExpenseRequest
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
