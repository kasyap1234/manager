package repository

import (
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	expenseRepo ExpenseRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		expenseRepo: &expenseRepository{db: db},
	}
}
