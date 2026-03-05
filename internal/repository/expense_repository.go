// Package repository
package repository

import (
	"context"

	"manager/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseRepository interface {
	GetExpenses() ([]*model.Expense, error)
	CreateExpense(expense *model.Expense) error
	UpdateExpense(expense *model.Expense) error
	DeleteExpense(id string) error
	GetExpenseByID(id string) (*model.Expense, error)
}

type expenseRepository struct {
	db *pgxpool.Pool
}

func NewExpenseRepository(db *pgxpool.Pool) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) GetExpenses() ([]*model.Expense, error) {
	query := `SELECT * FROM expenses`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expenses, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Expense])
	if err != nil {
		return nil, err
	}
	return &expenses, nil
	}

func (r *expenseRepository) CreateExpense(expense *model.Expense) error {
	query := `INSERT INTO expenses (amount, date, category, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.CreatedAt, expense.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *expenseRepository) UpdateExpense(expense *model.Expense) error {
	query := `UPDATE expenses SET amount = $1, date = $2, category = $3, description = $4, created_at = $5, updated_at = $6 WHERE id = $7`
	_, err := r.db.Exec(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.CreatedAt, expense.UpdatedAt, expense.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *expenseRepository) DeleteExpense(id string) error {
	query := `DELETE FROM expenses WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
