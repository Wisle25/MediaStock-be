package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// RatingRepository defines the interface for the rating repository
type RatingRepository interface {
	// CreateRating creates a new rating
	// Returning new rating id
	CreateRating(payload *entity.CreateRatingPayload) string

	// GetRatingsByAsset retrieves all ratings for a specific asset
	GetRatingsByAsset(assetId string) []entity.Rating

	// DeleteRating deletes a rating by its ID
	DeleteRating(id string)
}
