package use_case_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"testing"
	"time"
)

// Mocks for the dependencies
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(payload *entity.CreateTransactionPayload) string {
	args := m.Called(payload)
	return args.String(0)
}

func (m *MockTransactionRepository) GetTransactionByID(transactionID string) *entity.Transaction {
	args := m.Called(transactionID)
	return args.Get(0).(*entity.Transaction)
}

func (m *MockTransactionRepository) GetTransactionsByUser(userId string) []entity.PreviewTransaction {
	args := m.Called(userId)
	return args.Get(0).([]entity.PreviewTransaction)
}

func TestTransactionUseCase(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	transactionUseCase := use_case.NewTransactionUseCase(mockRepo)

	t.Run("ExecuteCreate", func(t *testing.T) {
		// Arrange
		payload := &entity.CreateTransactionPayload{
			UserID:      "user123",
			TotalAmount: 1000,
			AssetsId:    []string{"asset1", "asset2"},
		}
		expectedID := "transaction123"

		mockRepo.On("CreateTransaction", payload).Return(expectedID)

		// Action
		transactionID := transactionUseCase.ExecuteCreate(payload)

		// Assert
		assert.Equal(t, expectedID, transactionID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ExecuteGetByID", func(t *testing.T) {
		// Arrange
		transactionID := "transaction123"
		expectedTransaction := &entity.Transaction{
			ID:          "transaction123",
			TotalAmount: 1000,
			PurchasedAt: time.Now(),
			Items: []entity.PreviewTransactionItem{
				{ID: "item1", Title: "Asset 1", Price: "500"},
				{ID: "item2", Title: "Asset 2", Price: "500"},
			},
		}

		mockRepo.On("GetTransactionByID", transactionID).Return(expectedTransaction)

		// Action
		transaction := transactionUseCase.ExecuteGetByID(transactionID)

		// Assert
		assert.Equal(t, expectedTransaction, transaction)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ExecuteGetByUser", func(t *testing.T) {
		// Arrange
		userID := "user123"
		expectedTransactions := []entity.PreviewTransaction{
			{Id: "transaction1", TotalAmount: "1000", PurchasedAt: "2024-01-01T00:00:00Z"},
			{Id: "transaction2", TotalAmount: "2000", PurchasedAt: "2024-01-02T00:00:00Z"},
		}

		mockRepo.On("GetTransactionsByUser", userID).Return(expectedTransactions)

		// Action
		transactions := transactionUseCase.ExecuteGetByUser(userID)

		// Assert
		assert.Equal(t, expectedTransactions, transactions)
		mockRepo.AssertExpectations(t)
	})
}
