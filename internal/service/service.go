package service

type Service struct {
	expenseSvc *ExpenseService
}

func NewService(expenseSvc *ExpenseService) *Service {
	return &Service{
		expenseSvc: expenseSvc,
	}
}

func (s *Service) Expense() *ExpenseService {
	return s.expenseSvc
}
