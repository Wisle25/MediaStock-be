package use_case_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"testing"
)

// Mocks for the dependencies
type MockRatingRepository struct {
	mock.Mock
}

func (m *MockRatingRepository) CreateRating(payload *entity.CreateRatingPayload) string {
	args := m.Called(payload)
	return args.String(0)
}

func (m *MockRatingRepository) GetRatingsByAsset(assetId string) []entity.Rating {
	args := m.Called(assetId)
	return args.Get(0).([]entity.Rating)
}

func (m *MockRatingRepository) DeleteRating(id string) {
	m.Called(id)
}

type MockValidateRating struct {
	mock.Mock
}

func (m *MockValidateRating) ValidatePayload(payload *entity.CreateRatingPayload) {
	m.Called(payload)
}

func TestRatingUseCase(t *testing.T) {
	mockRepo := new(MockRatingRepository)
	mockValidator := new(MockValidateRating)
	ratingUseCase := use_case.NewRatingUseCase(mockRepo, mockValidator)

	t.Run("ExecuteCreate", func(t *testing.T) {
		// Arrange
		payload := &entity.CreateRatingPayload{
			AssetId:     "asset123",
			UserId:      "user123",
			Score:       5,
			Description: "Great asset!",
		}

		mockValidator.On("ValidatePayload", payload).Return(nil)
		mockRepo.On("CreateRating", payload).Return("rating123")

		// Action
		ratingId := ratingUseCase.ExecuteCreate(payload)

		// Assert
		assert.Equal(t, "rating123", ratingId)
		mockValidator.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ExecuteGetByAsset", func(t *testing.T) {
		// Arrange
		assetId := "asset123"
		expectedRatings := []entity.Rating{
			{ID: "rating1", Score: 5, Description: "Great asset!"},
			{ID: "rating2", Score: 4, Description: "Nice work!"},
		}

		mockRepo.On("GetRatingsByAsset", assetId).Return(expectedRatings)

		// Action
		ratings := ratingUseCase.ExecuteGetByAsset(assetId)

		// Assert
		assert.Equal(t, expectedRatings, ratings)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ExecuteDelete", func(t *testing.T) {
		// Arrange
		id := "rating123"

		mockRepo.On("DeleteRating", id).Return(nil)

		// Action
		ratingUseCase.ExecuteDelete(id)

		// Assert
		mockRepo.AssertExpectations(t)
	})
}
