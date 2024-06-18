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
func (u *CartUseCase) ExecuteAdd(payload *entity.CartPayload) {
	u.repository.AddToCart(payload)
}

// ExecuteGetAll returns carts asset by user id
func (u *CartUseCase) ExecuteGetAll(userId string) []entity.Cart {
	return u.repository.GetAllCarts(userId)
}

// ExecuteRemove remove cart from user
func (u *CartUseCase) ExecuteRemove(payload *entity.CartPayload) {
	u.repository.RemoveCart(payload)
}

// ExecuteRemoveAll Remove all carts item from user,
// WARNING! This only be called AFTER the checkout is success!
func (u *CartUseCase) ExecuteRemoveAll(userId string) {
	u.repository.RemoveAllCartByUser(userId)
}
