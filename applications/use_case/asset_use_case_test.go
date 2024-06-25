package use_case_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wisle25/media-stock-be/applications/file_statics"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/domains/entity"
	"mime/multipart"
	"path/filepath"
	"testing"
)

// Mocks for the dependencies
type MockAssetRepository struct {
	mock.Mock
}

func (m *MockAssetRepository) AddAsset(payload *entity.AssetPayload) string {
	args := m.Called(payload)
	return args.String(0)
}

func (m *MockAssetRepository) GetPreviewAssets(listCount int, pageList int, userId string) []entity.PreviewAsset {
	args := m.Called(listCount, pageList, userId)
	return args.Get(0).([]entity.PreviewAsset)
}

func (m *MockAssetRepository) GetDetailAsset(id string, userId string) *entity.Asset {
	args := m.Called(id, userId)
	return args.Get(0).(*entity.Asset)
}

func (m *MockAssetRepository) UpdateAsset(id string, payload *entity.AssetPayload) (string, string) {
	args := m.Called(id, payload)
	return args.String(0), args.String(1)
}

func (m *MockAssetRepository) DeleteAsset(id string) (string, string) {
	args := m.Called(id)
	return args.String(0), args.String(1)
}

func (m *MockAssetRepository) VerifyOwner(userId string, id string) {
	m.Called(userId, id)
}

func (m *MockAssetRepository) DownloadAsset(id string, userId string) (string, string) {
	args := m.Called(id, userId)
	return args.String(0), args.String(1)
}

type MockValidateAsset struct {
	mock.Mock
}

func (m *MockValidateAsset) ValidatePayload(payload *entity.AssetPayload) {
	m.Called(payload)
}

func TestAssetUseCase(t *testing.T) {
	mockAssetRepo := new(MockAssetRepository)
	mockFileProcessing := new(MockFileProcessing)
	mockFileUpload := new(MockFileUpload)
	mockValidation := new(MockValidateAsset)
	config := commons.LoadConfig("../..")

	assetUseCase := use_case.NewAssetUseCase(
		mockAssetRepo,
		mockFileProcessing,
		mockFileUpload,
		mockValidation,
		config,
	)

	t.Run("ExecuteAdd", func(t *testing.T) {
		payload := &entity.AssetPayload{
			File: &multipart.FileHeader{
				Filename: "test.png",
			},
		}

		mockValidation.On("ValidatePayload", payload).Return(nil)
		mockFileProcessing.On("CompressImage", mock.Anything, file_statics.WEBP).Return([]byte("compressed"), ".webp")
		mockFileProcessing.On("AddWatermark", mock.Anything).Return([]byte("watermarked"))
		mockFileUpload.On("UploadFile", mock.Anything, ".png").Return("original_link")
		mockFileUpload.On("UploadFile", mock.Anything, ".webp").Return("watermark_link")
		mockAssetRepo.On("AddAsset", payload).Return("asset_id")

		assetId := assetUseCase.ExecuteAdd(payload)

		assert.Equal(t, "asset_id", assetId)

		mockValidation.AssertExpectations(t)
		mockFileProcessing.AssertExpectations(t)
		mockFileUpload.AssertExpectations(t)
		mockAssetRepo.AssertExpectations(t)
	})

	t.Run("ExecuteGetAll", func(t *testing.T) {
		mockAssetRepo.On("GetPreviewAssets", 10, 1, "user1").Return([]entity.PreviewAsset{})

		assets := assetUseCase.ExecuteGetAll(10, 1, "user1")

		assert.NotNil(t, assets)
		mockAssetRepo.AssertExpectations(t)
	})

	t.Run("ExecuteGetDetail", func(t *testing.T) {
		mockAssetRepo.On("GetDetailAsset", "asset_id", "user1").Return(&entity.Asset{})

		asset := assetUseCase.ExecuteGetDetail("asset_id", "user1")

		assert.NotNil(t, asset)
		mockAssetRepo.AssertExpectations(t)
	})

	t.Run("ExecuteUpdate", func(t *testing.T) {
		payload := &entity.AssetPayload{
			File: &multipart.FileHeader{
				Filename: "test.png",
			},
		}

		mockValidation.On("ValidatePayload", payload).Return(nil)
		mockFileProcessing.On("CompressImage", mock.Anything, file_statics.WEBP).Return([]byte("compressed"), ".webp")
		mockFileProcessing.On("AddWatermark", mock.Anything).Return([]byte("watermarked"))
		mockFileUpload.On("UploadFile", mock.Anything, ".png").Return("original_link")
		mockFileUpload.On("UploadFile", mock.Anything, ".webp").Return("watermark_link")
		mockAssetRepo.On("UpdateAsset", "asset_id", payload).Return("old_original_link", "old_watermark_link")
		mockFileUpload.On("RemoveFile", "old_original_link").Return(nil)
		mockFileUpload.On("RemoveFile", "old_watermark_link").Return(nil)

		assetUseCase.ExecuteUpdate("asset_id", payload)

		mockValidation.AssertExpectations(t)
		mockFileProcessing.AssertExpectations(t)
		mockFileUpload.AssertExpectations(t)
		mockAssetRepo.AssertExpectations(t)
	})

	t.Run("ExecuteDelete", func(t *testing.T) {
		mockAssetRepo.On("DeleteAsset", "asset_id").Return("original_link", "watermark_link")
		mockFileUpload.On("RemoveFile", "original_link").Return(nil)
		mockFileUpload.On("RemoveFile", "watermark_link").Return(nil)

		assetUseCase.ExecuteDelete("asset_id")

		mockAssetRepo.AssertExpectations(t)
		mockFileUpload.AssertExpectations(t)
	})

	t.Run("ExecuteDownload", func(t *testing.T) {
		mockAssetRepo.On("DownloadAsset", "asset_id", "user1").Return("test", "file_path")
		mockFileUpload.On("GetFile", "file_path").Return([]byte("file_content"))

		filename, fileBuffer := assetUseCase.ExecuteDownload("asset_id", "user1")

		assert.Equal(t, "test"+filepath.Ext("file_path"), filename)
		assert.Equal(t, []byte("file_content"), fileBuffer)

		mockAssetRepo.AssertExpectations(t)
		mockFileUpload.AssertExpectations(t)
	})

	t.Run("VerifyAccess", func(t *testing.T) {
		mockAssetRepo.On("VerifyOwner", "user1", "asset_id").Return(nil)

		assetUseCase.VerifyAccess("user1", "asset_id")

		mockAssetRepo.AssertExpectations(t)
	})
}
