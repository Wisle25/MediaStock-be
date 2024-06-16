package use_case_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"io"
	"mime/multipart"
	"testing"
)

type MockAssetRepository struct {
	mock.Mock
}

func (m *MockAssetRepository) AddAsset(payload *entity.AddAssetPayload) string {
	args := m.Called(payload)

	return args.String(0)
}

func TestAssetUseCase(t *testing.T) {
	mockAssetRepository := new(MockAssetRepository)
	mockFileProcessing := new(MockFileProcessing)
	mockFileUpload := new(MockFileUpload)

	assetUseCase := use_case.NewAssetUseCase(mockAssetRepository, mockFileProcessing, mockFileUpload)

	t.Run("Execute Add", func(t *testing.T) {
		// Arrange
		payload := &entity.AddAssetPayload{
			Title:         "any title",
			File:          &multipart.FileHeader{},
			Description:   "any description",
			Details:       "any detail",
			OwnerId:       "user-123",
			OriginalLink:  "orilink",
			WatermarkLink: "waterlink",
		}

		file, _ := payload.File.Open()
		assetBuffer, _ := io.ReadAll(file)
		var compressBuffer []byte
		var watermarkBuffer []byte

		mockFileProcessing.On("CompressImage", assetBuffer, mock.Anything).Return(compressBuffer, ".webp")
		mockFileProcessing.On("AddWatermark", compressBuffer).Return(watermarkBuffer)
		mockFileUpload.On("UploadFile", assetBuffer, mock.Anything).Return(payload.OriginalLink).Once()
		mockFileUpload.On("UploadFile", watermarkBuffer, mock.Anything).Return(payload.WatermarkLink).Once()
		mockAssetRepository.On("AddAsset", payload).Return("asset-123")

		// Actions
		assetUseCase.ExecuteAdd(payload)

		// Assert
		mockFileUpload.AssertExpectations(t)
		mockAssetRepository.AssertExpectations(t)
		mockFileProcessing.AssertExpectations(t)
	})
}
