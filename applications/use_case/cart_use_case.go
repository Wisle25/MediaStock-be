package use_case

import (
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type CartUseCase struct {
	repository repository.CartRepository
}

func NewCartUseCase(repository repository.CartRepository) *CartUseCase {
	return &CartUseCase{
		repository,
	}
}

// ExecuteAdd will panic if user and asset is not existed
func (f *CartUseCase) ExecuteAdd(payload *entity.CartPayload) {
	f.repository.AddToCart(payload)
}

// ExecuteGetAll returns favorited asset by user id
func (f *CartUseCase) ExecuteGetAll(userId string) []entity.Cart {
	return f.repository.GetAllCarts(userId)
}

// ExecuteRemove remove favorite from user
func (f *CartUseCase) ExecuteRemove(payload *entity.CartPayload) {
	f.repository.RemoveCart(payload)
}
