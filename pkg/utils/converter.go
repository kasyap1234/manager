// Package utils
package utils

import "manager/internal/model"

func TransToExpense(transaction model.Transaction) (*model.Expense, error) {
	expense := &model.Expense{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Category:    transaction.Category,
		Description: transaction.Description,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
	return expense, nil
}

func ExpenseToTransaction(expense model.Expense) (*model.Transaction, error) {
	transaction := &model.Transaction{
		ID:          expense.ID,
		Amount:      expense.Amount,
		Date:        expense.Date,
		Category:    expense.Category,
		Description: expense.Description,
		CreatedAt:   expense.CreatedAt,
		UpdatedAt:   expense.UpdatedAt,
	}
	return transaction, nil
}

