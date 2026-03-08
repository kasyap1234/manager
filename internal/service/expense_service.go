package service

import (
	"manager/internal/model"
	"manager/internal/parser"
	"manager/internal/repository"
)

type ExpenseService struct {
	expenseRepo repository.ExpenseRepository
	parser      parser.Parser
}

func NewExpenseService(expenseRepository repository.ExpenseRepository, expenseParser parser.Parser) *ExpenseService {
	return &ExpenseService{expenseRepo: expenseRepository, parser: expenseParser}
}

func (s *ExpenseService) GetExpenses() ([]model.Expense, error) {
	return s.expenseRepo.GetExpenses()
}

func (s *ExpenseService) CreateExpense(sms string) (model.Expense, error) {
	expense, err := s.parser.Parse(sms)
	if err != nil {
		return model.Expense{}, err
	}

	return s.expenseRepo.CreateExpense(expense)
}

func (s *ExpenseService) UpdateExpense(id string, sms string) (model.Expense, error) {
	existingExpense, err := s.expenseRepo.GetExpenseByID(id)
	if err != nil {
		return model.Expense{}, err
	}

	expense, err := s.parser.Parse(sms)
	if err != nil {
		return model.Expense{}, err
	}
	expense.ID = existingExpense.ID
	expense.CreatedAt = existingExpense.CreatedAt

	return s.expenseRepo.UpdateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id string) error {
	return s.expenseRepo.DeleteExpense(id)
}

func (s *ExpenseService) GetExpenseByID(id string) (model.Expense, error) {
	return s.expenseRepo.GetExpenseByID(id)
}
