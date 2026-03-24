package repository

import (
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	transactionRepo TransactionRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		transactionRepo: NewTransactionRepository(db),
	}
}

func (r *Repository) Transaction() TransactionRepository {
	return r.transactionRepo
}
