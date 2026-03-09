// Package repository 
package repository

import (
	"context"
	"errors"

	"manager/internal/model"

	"github.com/jackc/pgx/v5"
)

var ErrExpenseNotFound = errors.New("expense not found")

type ExpenseRepository interface {
	GetExpenses() ([]model.Expense, error)
	CreateExpense(expense model.Expense) (model.Expense, error)
	UpdateExpense(expense model.Expense) (model.Expense, error)
	DeleteExpense(id string) error
	GetExpenseByID(id string) (model.Expense, error)
}

type expenseRepository struct {
	db *pgx.Conn
}

func NewExpenseRepository(db *pgx.Conn) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) GetExpenses() ([]model.Expense, error) {
	query := `SELECT id, amount, date, category, description, created_at, updated_at FROM expenses`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Expense])
}

func (r *expenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	query := `INSERT INTO expenses (amount, date, category, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, amount, date, category, description, created_at, updated_at`

	var createdExpense model.Expense
	err := r.db.QueryRow(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.CreatedAt, expense.UpdatedAt).Scan(
		&createdExpense.ID,
		&createdExpense.Amount,
		&createdExpense.Date,
		&createdExpense.Category,
		&createdExpense.Description,
		&createdExpense.CreatedAt,
		&createdExpense.UpdatedAt,
	)
	if err != nil {
		return model.Expense{}, err
	}

	return createdExpense, nil
}

func (r *expenseRepository) UpdateExpense(expense model.Expense) (model.Expense, error) {
	query := `UPDATE expenses SET amount = $1, date = $2, category = $3, description = $4, updated_at = $5 WHERE id = $6 RETURNING id, amount, date, category, description, created_at, updated_at`

	var updatedExpense model.Expense
	err := r.db.QueryRow(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.UpdatedAt, expense.ID).Scan(
		&updatedExpense.ID,
		&updatedExpense.Amount,
		&updatedExpense.Date,
		&updatedExpense.Category,
		&updatedExpense.Description,
		&updatedExpense.CreatedAt,
		&updatedExpense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Expense{}, ErrExpenseNotFound
		}
		return model.Expense{}, err
	}

	return updatedExpense, nil
}

func (r *expenseRepository) DeleteExpense(id string) error {
	query := `DELETE FROM expenses WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrExpenseNotFound
	}
	return nil
}

func (r *expenseRepository) GetExpenseByID(id string) (model.Expense, error) {
	query := `SELECT id, amount, date, category, description, created_at, updated_at FROM expenses WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)
	var expense model.Expense
	err := row.Scan(
		&expense.ID,
		&expense.Amount,
		&expense.Date,
		&expense.Category,
		&expense.Description,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Expense{}, ErrExpenseNotFound
		}
		return model.Expense{}, err
	}

	return expense, nil
}
