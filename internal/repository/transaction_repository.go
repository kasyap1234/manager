// Package repository
package repository

import (
	"context"
	"errors"
	"time"
	"manager/internal/model"

	"github.com/jackc/pgx/v5"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type TransactionRepository interface {
	GetTransactions() ([]model.Transaction, error)
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	UpdateTransaction(transaction model.Transaction) (model.Transaction, error)
	DeleteTransaction(id string) error
	GetTransactionByID(id string) (model.Transaction, error)
	GetTransactionsByCategory(category string) ([]model.Transaction, error)
	GetTransactionsByMerchant(merchant string) ([]model.Transaction, error)
	GetTransactionsByDate(date time.Time) ([]model.Transaction, error)
	GetTransactionsByMonth(year, month int) ([]model.Transaction, error)
	GetTransactionsByDateRange(start, end time.Time) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetTransactions() ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM transactions`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

	func (r *transactionRepository) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	query := `INSERT INTO transactions (amount, date, merchant,credit,category, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, amount, date, category, description, created_at, updated_at`

	var createdTransaction model.Transaction
	err := r.db.QueryRow(context.Background(), query, transaction.Amount, transaction.Date, transaction.Category, transaction.Description, transaction.CreatedAt, transaction.UpdatedAt).Scan(
		&createdTransaction.ID,
		&createdTransaction.Amount,
		&createdTransaction.Date,
		&createdTransaction.Merchant,
		&createdTransaction.Credit,
		&createdTransaction.Category,
		&createdTransaction.Description,
		&createdTransaction.CreatedAt,
		&createdTransaction.UpdatedAt,
	)
	if err != nil {
		return model.Transaction{}, err
	}

	return createdTransaction, nil
}

	func (r *transactionRepository) UpdateTransaction(transaction model.Transaction) (model.Transaction, error) {
	query := `UPDATE expenses SET amount = $1, date = $2, category = $3, description = $4, updated_at = $5 WHERE id = $6 RETURNING id, amount, date, category, description, created_at, updated_at`

	var updatedTransaction model.Transaction
	err := r.db.QueryRow(context.Background(), query, transaction.Amount, transaction.Date, transaction.Category, transaction.Description, transaction.UpdatedAt, transaction.ID).Scan(
		&updatedTransaction.ID,
		&updatedTransaction.Amount,
		&updatedTransaction.Date,
		&updatedTransaction.Merchant,
		&updatedTransaction.Credit,
		&updatedTransaction.Category,
		&updatedTransaction.Description,
		&updatedTransaction.CreatedAt,
		&updatedTransaction.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Transaction{}, ErrTransactionNotFound
		}
		return model.Transaction{}, err
	}

	return updatedTransaction, nil
}

func (r *transactionRepository) DeleteTransaction(id string) error {
	query := `DELETE FROM expenses WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrTransactionNotFound
	}
	return nil
}

func (r *transactionRepository) GetTransactionByID(id string) (model.Transaction, error) {
	query := `SELECT id, amount, date, category, description, created_at, updated_at FROM expenses WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)
	var transaction model.Transaction
	err := row.Scan(
		&transaction.ID,
		&transaction.Amount,
		&transaction.Date,
		&transaction.Merchant,
		&transaction.Credit,
		&transaction.Category,
		&transaction.Description,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Transaction{}, ErrTransactionNotFound
		}
		return model.Transaction{}, err
	}

	return transaction, nil
}


func(r*transactionRepository)GetTransactionsByCategory(category string) ([]model.Transaction, error) {	
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE category = $1`
	rows, err := r.db.Query(context.Background(), query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

func(r*transactionRepository)GetTransactionsByMerchant(merchant string) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE merchant = $1`
	rows, err := r.db.Query(context.Background(), query, merchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}

func(r*transactionRepository)GetTransactionsByDate(date time.Time) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE date = $1`
	rows, err := r.db.Query(context.Background(), query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}


func(r*transactionRepository)GetTransactionsByMonth(year , month int) ([]model.Transaction, error) {
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


func(r*transactionRepository)GetTransactionsByDateRange(start,end time.Time) ([]model.Transaction, error) {
	query := `SELECT id, amount, date, merchant, credit,category, description, created_at, updated_at FROM expenses WHERE date >= $1 AND date < $2`
	rows, err := r.db.Query(context.Background(), query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Transaction])
}


