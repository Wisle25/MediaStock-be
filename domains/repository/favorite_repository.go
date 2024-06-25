package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// FavoriteRepository defines the methods that any data storage provider needs to implement to manage favorite data.
type FavoriteRepository interface {
	// AddAsFavorite adds an asset to the user's favorites.
	// Should raise panic if either the asset or user does not exist.
	AddAsFavorite(payload *entity.FavoritePayload)

	// GetAllFavorites returns all favorited assets for a user by their user ID.
	GetAllFavorites(userId string) []entity.Favorite

	// RemoveFavorite removes an asset from the user's favorites.
	RemoveFavorite(payload *entity.FavoritePayload)
}
