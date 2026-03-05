package service


type Service struct {
	expenseService *ExpenseService
}

func NewService(expenseService *ExpenseService) *Service {
	return &Service{expenseService: expenseService}
}



