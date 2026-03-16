// Package repository
package repository

import (
	"context"
	"errors"
	"time"
	"manager/internal/model"

	"github.com/jackc/pgx/v5"
)

var ErrExpenseNotFound = errors.New("expense not found")

type ExpenseRepository interface {
	GetExpenses() ([]model.Transaction, error)
	CreateExpense(expense model.Transaction) (model.Transaction, error)
	UpdateExpense(expense model.Transaction) (model.Transaction, error)
	DeleteExpense(id string) error
	GetExpenseByID(id string) (model.Transaction, error)
	GetExpensesByCategory(category string) ([]model.Transaction, error)
	GetExpensesByMerchant(merchant string) ([]model.Transaction, error)
	GetExpensesByDate(date time.Time) ([]model.Transaction, error)
	GetExpensesByMonth(year, month int) ([]model.Transaction, error)
	GetExpensesByDateRange(start, end time.Time) ([]model.Transaction, error)
}

type expenseRepository struct {
	db *pgx.Conn
}

func NewExpenseRepository(db *pgx.Conn) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) GetExpenses() ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

func (r *expenseRepository) CreateExpense(expense model.Transaction) (model.Transaction, error) {
	query := `INSERT INTO expenses (amount, date, merchant,credit,category, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, amount, date, category, description, created_at, updated_at`

	var createdExpense model.Transaction
	err := r.db.QueryRow(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.CreatedAt, expense.UpdatedAt).Scan(
		&createdExpense.ID,
		&createdExpense.Amount,
		&createdExpense.Date,
		&createdExpense.Merchant,
		&createdExpense.Credit,
		&createdExpense.Category,
		&createdExpense.Description,
		&createdExpense.CreatedAt,
		&createdExpense.UpdatedAt,
	)
	if err != nil {
		return model.Transaction{}, err
	}

	return createdExpense, nil
}

func (r *expenseRepository) UpdateExpense(expense model.Transaction) (model.Transaction, error) {
	query := `UPDATE expenses SET amount = $1, date = $2, category = $3, description = $4, updated_at = $5 WHERE id = $6 RETURNING id, amount, date, category, description, created_at, updated_at`

	var updatedExpense model.Transaction
	err := r.db.QueryRow(context.Background(), query, expense.Amount, expense.Date, expense.Category, expense.Description, expense.UpdatedAt, expense.ID).Scan(
		&updatedExpense.ID,
		&updatedExpense.Amount,
		&updatedExpense.Date,
		&updatedExpense.Merchant,
		&updatedExpense.Credit,
		&updatedExpense.Category,
		&updatedExpense.Description,
		&updatedExpense.CreatedAt,
		&updatedExpense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Transaction{}, ErrExpenseNotFound
		}
		return model.Transaction{}, err
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

func (r *expenseRepository) GetExpenseByID(id string) (model.Transaction, error) {
	query := `SELECT id, amount, date, category, description, created_at, updated_at FROM expenses WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)
	var expense model.Transaction
	err := row.Scan(
		&expense.ID,
		&expense.Amount,
		&expense.Date,
		&expense.Merchant,
		&expense.Credit,
		&expense.Category,
		&expense.Description,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Transaction{}, ErrExpenseNotFound
		}
		return model.Transaction{}, err
	}

	return expense, nil
}


func(r*expenseRepository)GetExpensesByCategory(category string) ([]model.Transaction, error) {	
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE category = $1`
	rows, err := r.db.Query(context.Background(), query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

func(r*expenseRepository)GetExpensesByMerchant(merchant string) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE merchant = $1`
	rows, err := r.db.Query(context.Background(), query, merchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

func(r*expenseRepository)GetExpensesByDate(date time.Time) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE date = $1`
	rows, err := r.db.Query(context.Background(), query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}


func(r*expenseRepository)GetExpensesByMonth(year , month int) ([]model.Transaction, error) {
    start :=time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end :=start.AddDate(0, 1, 0)
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE date >= $1 AND date < $2`
	rows, err := r.db.Query(context.Background(), query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}


func(r*expenseRepository)GetExpensesByDateRange(start,end time.Time) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE date >= $1 AND date < $2`
	rows, err := r.db.Query(context.Background(), query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}


