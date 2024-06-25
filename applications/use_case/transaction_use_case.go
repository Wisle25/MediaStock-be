package use_case

import (
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

// TransactionUseCase handles the business logic for transaction operations.
type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

// NewTransactionUseCase creates a new instance of TransactionUseCase.
func NewTransactionUseCase(
	transactionRepository repository.TransactionRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepository,
	}
}

// ExecuteCreate creates a new transaction and returns the transaction ID.
func (uc *TransactionUseCase) ExecuteCreate(payload *entity.CreateTransactionPayload) string {
	id := uc.transactionRepository.CreateTransaction(payload)
	return id
}

// ExecuteGetByID retrieves a transaction by its ID.
func (uc *TransactionUseCase) ExecuteGetByID(transactionID string) *entity.Transaction {
	return uc.transactionRepository.GetTransactionByID(transactionID)
}

// ExecuteGetByUser retrieves all transactions for a specific user.
func (uc *TransactionUseCase) ExecuteGetByUser(userId string) []entity.PreviewTransaction {
	return uc.transactionRepository.GetTransactionsByUser(userId)
}
