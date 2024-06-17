package repository

import "github.com/wisle25/media-stock-be/domains/entity"

type AssetRepository interface {
	// AddAsset Adding to database.
	// Returning the created asset's ID.
	AddAsset(payload *entity.AddAssetPayload) string

	// GetPreviewAssets Getting all assets but only the preview
	GetPreviewAssets(listCount int, pageList int) []entity.PreviewAsset

	// GetDetailAsset detailed of asset
	GetDetailAsset(id string, userId string) *entity.Asset

	// UpdateAsset Updating
	// Returning old both original and watermarked asset to be removed.
	UpdateAsset(id string, payload *entity.AddAssetPayload) (string, string)

	// DeleteAsset Deleting.
	// Returning old both original and watermarked asset to be removed.
	DeleteAsset(id string) (string, string)

	// VerifyOwner verifying the owner
	// Should raise panic if userId is not the owner
	VerifyOwner(userId string, id string)
}
