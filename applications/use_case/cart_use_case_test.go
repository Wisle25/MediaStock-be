package use_case_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

// Mocks for the dependencies
type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) AddToCart(payload *entity.CartPayload) {
	m.Called(payload)
}

func (m *MockCartRepository) GetAllCarts(userId string) []entity.Cart {
	args := m.Called(userId)
	return args.Get(0).([]entity.Cart)
}

func (m *MockCartRepository) RemoveCart(payload *entity.CartPayload) {
	m.Called(payload)
}

func (m *MockCartRepository) RemoveAllCartByUser(userId string) {
	m.Called(userId)
}

func TestCartUseCase(t *testing.T) {
	mockRepo := new(MockCartRepository)
	cartUseCase := use_case.NewCartUseCase(mockRepo)

	t.Run("Execute Add", func(t *testing.T) {
		// Arrange
		payload := &entity.CartPayload{
			UserId:  "user123",
			AssetId: "asset123",
		}

		mockRepo.On("AddToCart", payload).Return(nil)

		// Action
		cartUseCase.ExecuteAdd(payload)

		// Assert
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute GetAll", func(t *testing.T) {
		// Arrange
		userId := "user123"
		expectedCarts := []entity.Cart{
			{Id: "1", Title: "Asset 1", Price: "100", FilePath: "/path/to/asset1"},
			{Id: "2", Title: "Asset 2", Price: "200", FilePath: "/path/to/asset2"},
		}

		mockRepo.On("GetAllCarts", userId).Return(expectedCarts)

		// Action
		carts := cartUseCase.ExecuteGetAll(userId)

		// Assert
		assert.Equal(t, expectedCarts, carts)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute Remove", func(t *testing.T) {
		// Arrange
		payload := &entity.CartPayload{
			UserId:  "user123",
			AssetId: "asset123",
		}

		mockRepo.On("RemoveCart", payload).Return(nil)

		// Action
		cartUseCase.ExecuteRemove(payload)

		// Assert
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute RemoveAll", func(t *testing.T) {
		// Arrange
		userId := "user123"

		mockRepo.On("RemoveAllCartByUser", userId).Return(nil)

		// Action
		cartUseCase.ExecuteRemoveAll(userId)

		// Assert
		mockRepo.AssertExpectations(t)
	})
}
