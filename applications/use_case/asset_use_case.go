package use_case

import (
	"github.com/wisle25/media-stock-be/applications/file_statics"
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"io"
	"path/filepath"
)

type AssetUseCase struct {
	assetRepository repository.AssetRepository
	fileProcessing  file_statics.FileProcessing
	fileUpload      file_statics.FileUpload
	validation      validation.ValidateAsset
}

func NewAssetUseCase(
	assetRepository repository.AssetRepository,
	processing file_statics.FileProcessing,
	upload file_statics.FileUpload,
	validation validation.ValidateAsset,
) *AssetUseCase {
	return &AssetUseCase{
		assetRepository,
		processing,
		upload,
		validation,
	}
}

func (uc *AssetUseCase) ExecuteAdd(payload *entity.AddAssetPayload) string {
	//uc.validation.ValidatePayload(payload)
	uc.handleFileAsset(payload)

	// Finally add it to repository
	return uc.assetRepository.AddAsset(payload)
}

func (uc *AssetUseCase) ExecuteGetAll(listCount int, pageList int) []entity.PreviewAsset {
	return uc.assetRepository.GetPreviewAssets(listCount, pageList)
}

func (uc *AssetUseCase) ExecuteGetDetail(id string, userId string) *entity.Asset {
	return uc.assetRepository.GetDetailAsset(id, userId)
}

func (uc *AssetUseCase) ExecuteUpdate(id string, payload *entity.AddAssetPayload) {
	//uc.validation.ValidatePayload(payload)
	uc.handleFileAsset(payload)
	oldOriginalAsset, oldWatermarkedAsset := uc.assetRepository.UpdateAsset(id, payload)

	// Then remove the old assets from storage
	uc.fileUpload.RemoveFile(oldOriginalAsset)
	uc.fileUpload.RemoveFile(oldWatermarkedAsset)
}

func (uc *AssetUseCase) ExecuteDelete(id string) {
	// Delete from repository
	originalAsset, watermarkedAsset := uc.assetRepository.DeleteAsset(id)

	// Remove file asset from storage
	uc.fileUpload.RemoveFile(originalAsset)
	uc.fileUpload.RemoveFile(watermarkedAsset)
}

func (uc *AssetUseCase) VerifyAccess(userId string, id string) {
	// LATER, Verify if he is admin later

	// Verify Owner
	uc.assetRepository.VerifyOwner(userId, id)
}

func (uc *AssetUseCase) handleFileAsset(payload *entity.AddAssetPayload) {
	// Open asset file
	file, _ := payload.File.Open()
	assetBuffer, _ := io.ReadAll(file)

	// Make watermarked version of the original asset
	compressBuffer, extension := uc.fileProcessing.CompressImage(assetBuffer, file_statics.WEBP)
	watermarkBuffer := uc.fileProcessing.AddWatermark(compressBuffer)

	// Upload both watermarked and original asset
	payload.OriginalLink = uc.fileUpload.UploadFile(assetBuffer, filepath.Ext(payload.File.Filename))
	payload.WatermarkLink = uc.fileUpload.UploadFile(watermarkBuffer, extension)
}
