package use_case

import (
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

// FavoriteUseCase handles the business logic for favorite operations.
type FavoriteUseCase struct {
	repository repository.FavoriteRepository
}

// NewFavoriteUseCase creates a new instance of FavoriteUseCase.
func NewFavoriteUseCase(repository repository.FavoriteRepository) *FavoriteUseCase {
	return &FavoriteUseCase{
		repository,
	}
}

// ExecuteAdd adds an asset to the user's favorites.
// Will panic if the user or asset does not exist.
func (f *FavoriteUseCase) ExecuteAdd(payload *entity.FavoritePayload) {
	f.repository.AddAsFavorite(payload)
}

// ExecuteGetAll retrieves all favorited assets for a user by user ID.
func (f *FavoriteUseCase) ExecuteGetAll(userId string) []entity.Favorite {
	return f.repository.GetAllFavorites(userId)
}

// ExecuteRemove removes an asset from the user's favorites.
func (f *FavoriteUseCase) ExecuteRemove(payload *entity.FavoritePayload) {
	f.repository.RemoveFavorite(payload)
}
