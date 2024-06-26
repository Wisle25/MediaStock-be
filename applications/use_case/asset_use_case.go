package use_case

import (
	"github.com/wisle25/media-stock-be/applications/file_statics"
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"io"
	"path"
	"path/filepath"
)

// AssetUseCase is the use case implementation for managing assets.
type AssetUseCase struct {
	assetRepository repository.AssetRepository
	fileProcessing  file_statics.FileProcessing
	fileUpload      file_statics.FileUpload
	validation      validation.ValidateAsset
	config          *commons.Config
}

// NewAssetUseCase creates a new instance of AssetUseCase.
func NewAssetUseCase(
	assetRepository repository.AssetRepository,
	processing file_statics.FileProcessing,
	upload file_statics.FileUpload,
	validation validation.ValidateAsset,
	config *commons.Config,
) *AssetUseCase {
	return &AssetUseCase{
		assetRepository,
		processing,
		upload,
		validation,
		config,
	}
}

// ExecuteAdd validates the payload, processes the file asset, and adds the asset to the repository.
func (uc *AssetUseCase) ExecuteAdd(payload *entity.AssetPayload) string {
	uc.validation.ValidatePayload(payload)
	uc.handleFileAsset(payload)

	// Adjust the link to be added with Minio
	minioUrl := uc.config.MinioUrl + uc.config.MinioBucket + "/"
	payload.OriginalLink = minioUrl + payload.OriginalLink
	payload.WatermarkLink = minioUrl + payload.WatermarkLink

	// ...
	return uc.assetRepository.AddAsset(payload)
}

// ExecuteGetAll retrieves a list of preview assets for the given user.
func (uc *AssetUseCase) ExecuteGetAll(listCount int, pageList int, userId string, sortBy string, search string, category string) []entity.PreviewAsset {
	return uc.assetRepository.GetPreviewAssets(listCount, pageList, userId, sortBy, search, category)
}

// ExecuteGetPurchased Get the purchased item by the user
func (uc *AssetUseCase) ExecuteGetPurchased(userId string) []entity.SimpleAsset {
	return uc.assetRepository.GetPurchasedAsset(userId)
}

// ExecuteGetOwned Owned/sold item by the user
func (uc *AssetUseCase) ExecuteGetOwned(userId string) []entity.SimpleAsset {
	return uc.assetRepository.GetOwnedAsset(userId)
}

// ExecuteGetDetail retrieves the details of a specific asset.
func (uc *AssetUseCase) ExecuteGetDetail(id string, userId string) *entity.Asset {
	return uc.assetRepository.GetDetailAsset(id, userId)
}

// ExecuteUpdate validates the payload, processes the file asset, updates the asset, and removes old assets.
func (uc *AssetUseCase) ExecuteUpdate(id string, payload *entity.AssetPayload) {
	uc.validation.ValidatePayload(payload)
	uc.handleFileAsset(payload)

	// Update the asset in the repository and get the old asset links
	oldOriginalAsset, oldWatermarkedAsset := uc.assetRepository.UpdateAsset(id, payload)

	// Remove the old assets from storage
	uc.fileUpload.RemoveFile(path.Base(oldOriginalAsset))
	uc.fileUpload.RemoveFile(path.Base(oldWatermarkedAsset))
}

// ExecuteDelete deletes the asset from the repository and removes the file assets from storage.
func (uc *AssetUseCase) ExecuteDelete(id string) {
	// Delete the asset from the repository and get the asset links
	originalAsset, watermarkedAsset := uc.assetRepository.DeleteAsset(id)

	// Remove the file assets from storage
	uc.fileUpload.RemoveFile(path.Base(originalAsset))
	uc.fileUpload.RemoveFile(path.Base(watermarkedAsset))
}

// VerifyAccess verifies if the user has access to the asset.
func (uc *AssetUseCase) VerifyAccess(userId string, id string) {
	// LATER, Verify if the user is an admin later
	uc.assetRepository.VerifyOwner(userId, id)
}

// ExecuteDownload returns the asset link and the file buffer to be downloaded.
func (uc *AssetUseCase) ExecuteDownload(id string, userId string) (string, []byte) {
	title, fileLink := uc.assetRepository.DownloadAsset(id, userId)
	fileBuffer := uc.fileUpload.GetFile(path.Base(fileLink))

	return title + filepath.Ext(fileLink), fileBuffer
}

// handleFileAsset processes the file asset by compressing and adding a watermark, then uploads it.
func (uc *AssetUseCase) handleFileAsset(payload *entity.AssetPayload) {
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
