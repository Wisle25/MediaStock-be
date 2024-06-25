package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// TransactionRepository defines the methods that any data storage provider needs to implement to manage transaction data.
type TransactionRepository interface {
	// CreateTransaction creates a new transaction and returns the transaction ID.
	CreateTransaction(payload *entity.CreateTransactionPayload) string

	// GetTransactionByID retrieves a transaction by its ID.
	GetTransactionByID(transactionID string) *entity.Transaction

	// GetTransactionsByUser retrieves all transactions for a specific user by their user ID.
	GetTransactionsByUser(userID string) []entity.PreviewTransaction
}
