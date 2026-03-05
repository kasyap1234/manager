package service

import (
	"fmt"
	"manager/internal/model"
	"manager/internal/parser"
	"manager/internal/repository"
	"manager/pkg/utils"
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

func (s *ExpenseService) CreateExpense(sms string) error {
	transaction, err := parser.SMSParser.Parse(sms)
	expense,err := utils.TransToExpense(transaction)
	if err != nil {
		return err
	}
	return s.expenseRepository.CreateExpense(expense)
}

func (s *ExpenseService) UpdateExpense(sms string) error {
     transaction, err := parser.SMSParser.Parse(sms)
	 expense,err := utils.TransToExpense(transaction)
	 if err != nil {
		return err
	 }
	 return s.expenseRepository.UpdateExpense(expense)
}

func (s *ExpenseService) DeleteExpense(id string) error {
	expense, err := s.expenseRepository.GetExpenseByID(id)
	if err != nil {
		return err
	}
	if expense.ID == "" {
		return fmt.Errorf("expense not found")
	}
	return s.expenseRepository.DeleteExpense(expense.ID)
}

// GetExpenseByID is a method that returns a single expense by its ID
func (s *ExpenseService) GetExpenseByID(id string) (model.Expense, error) {
	expense, err := s.expenseRepository.GetExpenseByID(id)
	if err != nil {
		return model.Expense{}, err
	}
	if expense.ID == "" {
		return model.Expense{}, fmt.Errorf("expense not found")
	}
	return expense, nil
}
