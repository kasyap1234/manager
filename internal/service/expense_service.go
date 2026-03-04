package service

import (
	"manager/internal/model"
	"manager/internal/repository"
)

type ExpenseService struct {
	expenseRepository repository.ExpenseRepository
}



func NewExpenseService(expenseRepository repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{expenseRepository: expenseRepository}
}

func (s *ExpenseService) GetExpenses() ([]model.Expense, error) {
	return s.expenseRepository.GetExpenses()
}

func (s *ExpenseService) CreateExpense(expense model.Expense) error {
	return s.expenseRepository.CreateExpense(expense)
}

func (s *ExpenseService) UpdateExpense(expense model.Expense) error {
	return s.expenseRepository.UpdateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id string) error {
	return s.expenseRepository.DeleteExpense(id)
}

// GetExpenseById is a method that returns a single expense by its ID
func (s *ExpenseService) GetExpenseById(id string) (model.Expense, error) {
	return s.expenseRepository.GetExpenseById(id)
}
