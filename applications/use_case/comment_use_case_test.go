package use_case_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"testing"
)

// Mocks for the dependencies
type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) CreateComment(payload *entity.CreateCommentPayload) string {
	args := m.Called(payload)
	return args.String(0)
}

func (m *MockCommentRepository) GetCommentsByAsset(assetId string) []entity.Comment {
	args := m.Called(assetId)
	return args.Get(0).([]entity.Comment)
}

func (m *MockCommentRepository) UpdateComment(id string, content string) {
	m.Called(id, content)
}

func (m *MockCommentRepository) DeleteComment(id string) {
	m.Called(id)
}

type MockValidateComment struct {
	mock.Mock
}

func (m *MockValidateComment) ValidatePayload(payload *entity.CreateCommentPayload) {
	m.Called(payload)
}

func (m *MockValidateComment) ValidateUpdate(payload *entity.EditCommentPayload) {
	m.Called(payload)
}

func TestCommentUseCase(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	mockValidator := new(MockValidateComment)
	commentUseCase := use_case.NewCommentUseCase(mockRepo, mockValidator)

	t.Run("Execute Create", func(t *testing.T) {
		// Arrange
		payload := &entity.CreateCommentPayload{
			AssetId: "asset123",
			UserId:  "user123",
			Content: "Great asset!",
		}

		mockValidator.On("ValidatePayload", payload).Return(nil)
		mockRepo.On("CreateComment", payload).Return("comment123")

		// Action
		commentId := commentUseCase.ExecuteCreate(payload)

		// Assert
		assert.Equal(t, "comment123", commentId)
		mockValidator.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute GetByAsset", func(t *testing.T) {
		// Arrange
		assetId := "asset123"
		expectedComments := []entity.Comment{
			{Id: "comment1", Content: "Great asset!"},
			{Id: "comment2", Content: "Nice work!"},
		}

		mockRepo.On("GetCommentsByAsset", assetId).Return(expectedComments)

		// Action
		comments := commentUseCase.ExecuteGetByAsset(assetId)

		// Assert
		assert.Equal(t, expectedComments, comments)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute Update", func(t *testing.T) {
		// Arrange
		id := "comment123"
		payload := &entity.EditCommentPayload{
			Content: "Updated comment",
		}

		mockValidator.On("ValidateUpdate", payload).Return(nil)
		mockRepo.On("UpdateComment", id, payload.Content).Return(nil)

		// Action
		commentUseCase.ExecuteUpdate(id, payload)

		// Assert
		mockValidator.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Execute Delete", func(t *testing.T) {
		// Arrange
		id := "comment123"

		mockRepo.On("DeleteComment", id).Return(nil)

		// Action
		commentUseCase.ExecuteDelete(id)

		// Assert
		mockRepo.AssertExpectations(t)
	})
}
