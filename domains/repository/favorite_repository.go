package repository

import "github.com/wisle25/media-stock-be/domains/entity"

type FavoriteRepository interface {
	// AddAsFavorite should raise panic if both asset or user is not existed
	AddAsFavorite(payload *entity.FavoritePayload)

	// GetAllFavorites returns favorited asset by user id
	GetAllFavorites(userId string) []entity.Favorite

	// RemoveFavorite remove favorite from user
	RemoveFavorite(payload *entity.FavoritePayload)
}
