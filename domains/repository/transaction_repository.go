package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// TransactionRepository defines the interface for transaction repository
type TransactionRepository interface {
	// CreateTransaction creates a new transaction and returns the transaction entity.
	// Returns transaction ID.
	CreateTransaction(payload *entity.CreateTransactionPayload) string

	// GetTransactionByID retrieves a transaction by its ID
	GetTransactionByID(transactionID string) *entity.Transaction

	// GetTransactionsByUser retrieves all transactions for a specific user
	GetTransactionsByUser(userID string) []entity.PreviewTransaction
}
