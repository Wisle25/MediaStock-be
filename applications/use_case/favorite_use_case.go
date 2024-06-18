package use_case

import (
	"fmt"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type FavoriteUseCase struct {
	repository repository.FavoriteRepository
}

func NewFavoriteUseCase(repository repository.FavoriteRepository) *FavoriteUseCase {
	return &FavoriteUseCase{
		repository,
	}
}

// ExecuteAdd will panic if user and asset is not existed
func (f *FavoriteUseCase) ExecuteAdd(payload *entity.FavoritePayload) {
	fmt.Printf("assetId: %s, userId: %s", payload.AssetId, payload.UserId)
	f.repository.AddAsFavorite(payload)
}

// ExecuteGetAll returns favorited asset by user id
func (f *FavoriteUseCase) ExecuteGetAll(userId string) []entity.Favorite {
	return f.repository.GetAllFavorites(userId)
}

// ExecuteRemove remove favorite from user
func (f *FavoriteUseCase) ExecuteRemove(payload *entity.FavoritePayload) {
	f.repository.RemoveFavorite(payload)
}
