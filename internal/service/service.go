// Package service
package service

type Service struct {
	transactionSvc *TransactionService
}

func NewService(transactionSvc *TransactionService) *Service {
	return &Service{
		transactionSvc: transactionSvc,
	}
}

func (s *Service) Transaction() *TransactionService {
	return s.transactionSvc
}
