package use_case

import (
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

// CartUseCase handles the business logic for cart operations.
type CartUseCase struct {
	repository repository.CartRepository
}

// NewCartUseCase creates a new instance of CartUseCase.
func NewCartUseCase(repository repository.CartRepository) *CartUseCase {
	return &CartUseCase{
		repository,
	}
}

// ExecuteAdd adds an item to the cart.
// Will panic if the user or asset does not exist.
func (u *CartUseCase) ExecuteAdd(payload *entity.CartPayload) {
	u.repository.AddToCart(payload)
}

// ExecuteGetAll retrieves all cart items for a user by user ID.
func (u *CartUseCase) ExecuteGetAll(userId string) []entity.Cart {
	return u.repository.GetAllCarts(userId)
}

// ExecuteRemove removes an item from the cart.
func (u *CartUseCase) ExecuteRemove(payload *entity.CartPayload) {
	u.repository.RemoveCart(payload)
}

// ExecuteRemoveAll removes all items from the cart for a user.
// WARNING: This should only be called AFTER a successful checkout.
func (u *CartUseCase) ExecuteRemoveAll(userId string) {
	u.repository.RemoveAllCartByUser(userId)
}
