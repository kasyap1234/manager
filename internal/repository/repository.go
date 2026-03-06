package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	expenseRepo ExpenseRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		expenseRepo: &expenseRepository{db: db},
	}
}
