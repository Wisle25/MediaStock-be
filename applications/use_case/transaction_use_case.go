package use_case

import (
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUseCase(
	transactionRepository repository.TransactionRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepository,
	}
}

func (uc *TransactionUseCase) ExecuteCreate(payload *entity.CreateTransactionPayload) string {
	id := uc.transactionRepository.CreateTransaction(payload)

	return id
}

func (uc *TransactionUseCase) ExecuteGetByID(transactionID string) *entity.Transaction {
	return uc.transactionRepository.GetTransactionByID(transactionID)
}

func (uc *TransactionUseCase) ExecuteGetByUser(userId string) []entity.PreviewTransaction {
	return uc.transactionRepository.GetTransactionsByUser(userId)
}
