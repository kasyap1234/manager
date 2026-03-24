package service

import (
	"time"

	"manager/internal/model"
	"manager/internal/parser"
	"manager/internal/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
	parser          parser.Parser
}

func NewTransactionService(transactionRepository repository.TransactionRepository, transactionParser parser.Parser) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepository,
		parser:          transactionParser,
	}
}

func (s *TransactionService) GetTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactions()
}

func (s *TransactionService) CreateTransaction(sms string) (model.Transaction, error) {
	transaction, err := s.parser.Parse(sms)
	if err != nil {
		return model.Transaction{}, err
	}

	return s.transactionRepo.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateTransaction(id string, sms string) (model.Transaction, error) {
	existingTransaction, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction, err := s.parser.Parse(sms)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction.ID = existingTransaction.ID
	transaction.CreatedAt = existingTransaction.CreatedAt
	transaction.UpdatedAt = time.Now()

	return s.transactionRepo.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(id string) error {
	return s.transactionRepo.DeleteTransaction(id)
}

func (s *TransactionService) GetTransactionByID(id string) (model.Transaction, error) {
	return s.transactionRepo.GetTransactionByID(id)
}

func (s *TransactionService) GetTransactionsByCategory(category string) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByCategory(category)
}

func (s *TransactionService) GetTransactionsByMerchant(merchant string) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByMerchant(merchant)
}

func (s *TransactionService) GetTransactionsByDate(date time.Time) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByDate(date)
}

func (s *TransactionService) GetTransactionsByMonth(year, month int) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByMonth(year, month)
}

func (s *TransactionService) GetTransactionsByDateRange(start, end time.Time) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByDateRange(start, end)
}
