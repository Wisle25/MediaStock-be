package use_case_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"testing"
)

// Mocks for the dependencies
type MockFavoriteRepository struct {
	mock.Mock
}

func (m *MockFavoriteRepository) AddAsFavorite(payload *entity.FavoritePayload) {
	m.Called(payload)
}

func (m *MockFavoriteRepository) GetAllFavorites(userId string) []entity.Favorite {
	args := m.Called(userId)
	return args.Get(0).([]entity.Favorite)
}

func (m *MockFavoriteRepository) RemoveFavorite(payload *entity.FavoritePayload) {
	m.Called(payload)
}

func TestFavoriteUseCase(t *testing.T) {
	mockRepo := new(MockFavoriteRepository)
	favoriteUseCase := use_case.NewFavoriteUseCase(mockRepo)

	t.Run("Execute Add", func(t *testing.T) {
		// Arrange
		payload := &entity.FavoritePayload{
			UserId:  "user123",
			AssetId: "asset123",
		}

		mockRepo.On("AddAsFavorite", payload).Return(nil)

		// Action
		favoriteUseCase.ExecuteAdd(payload)

		// Assert
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute GetAll", func(t *testing.T) {
		// Arrange
		userId := "user123"
		expectedFavorites := []entity.Favorite{
			{Id: "favorite1", Title: "Asset 1", Price: "10", FilePath: "path/to/asset1"},
			{Id: "favorite2", Title: "Asset 2", Price: "20", FilePath: "path/to/asset2"},
		}

		mockRepo.On("GetAllFavorites", userId).Return(expectedFavorites)

		// Action
		favorites := favoriteUseCase.ExecuteGetAll(userId)

		// Assert
		assert.Equal(t, expectedFavorites, favorites)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute Remove", func(t *testing.T) {
		// Arrange
		payload := &entity.FavoritePayload{
			UserId:  "user123",
			AssetId: "asset123",
		}

		mockRepo.On("RemoveFavorite", payload).Return(nil)

		// Action
		favoriteUseCase.ExecuteRemove(payload)

		// Assert
		mockRepo.AssertExpectations(t)
	})
}
