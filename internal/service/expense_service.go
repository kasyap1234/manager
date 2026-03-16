package service

import (
	"manager/internal/model"
	"manager/internal/parser"
	"manager/internal/repository"
	"time"
)

type ExpenseService struct {
	expenseRepo repository.ExpenseRepository
	parser      parser.Parser
}

func NewExpenseService(expenseRepository repository.ExpenseRepository, expenseParser parser.Parser) *ExpenseService {
	return &ExpenseService{expenseRepo: expenseRepository, parser: expenseParser}
}

func (s *ExpenseService) GetExpenses() ([]model.Transaction, error) {
	return s.expenseRepo.GetExpenses()
}

func (s *ExpenseService) CreateExpense(sms string) (model.Transaction, error) {
	expense, err := s.parser.Parse(sms)
	if err != nil {
		return model.Transaction{}, err
	}

	return s.expenseRepo.CreateExpense(expense)
}

func (s *ExpenseService) UpdateExpense(id string, sms string) (model.Transaction, error) {
	existingExpense, err := s.expenseRepo.GetExpenseByID(id)
	if err != nil {
		return model.Transaction{}, err
	}

	expense, err := s.parser.Parse(sms)
	if err != nil {
		return model.Transaction{}, err
	}
	expense.ID = existingExpense.ID
	expense.CreatedAt = existingExpense.CreatedAt

	return s.expenseRepo.UpdateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id string) error {
	return s.expenseRepo.DeleteExpense(id)
}

func (s *ExpenseService) GetExpenseByID(id string) (model.Transaction, error) {
	return s.expenseRepo.GetExpenseByID(id)
}


func (s *ExpenseService) GetExpensesByCategory(category string) ([]model.Transaction, error) {
	return s.expenseRepo.GetExpensesByCategory(category)
}

func (s *ExpenseService) GetExpensesByMerchant(merchant string) ([]model.Transaction, error) {
	return s.expenseRepo.GetExpensesByMerchant(merchant)
}

func (s *ExpenseService) GetExpensesByDate(date time.Time) ([]model.Transaction, error) {
	return s.expenseRepo.GetExpensesByDate(date)
}


func (s *ExpenseService) GetExpensesByMonth(year, month int) ([]model.Transaction, error) {
	return s.expenseRepo.GetExpensesByMonth(year, month)
}

func (s *ExpenseService) GetExpensesByDateRange(start, end time.Time) ([]model.Transaction, error) {
	return s.expenseRepo.GetExpensesByDateRange(start, end)
}
