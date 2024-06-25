package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// CartRepository defines the methods that any data storage provider needs to implement to manage cart data.
type CartRepository interface {
	// AddToCart adds an item to the cart.
	// Should raise panic if either the asset or user does not exist.
	AddToCart(payload *entity.CartPayload)

	// GetAllCarts returns all cart items for a user by their user ID.
	GetAllCarts(userId string) []entity.Cart

	// RemoveCart removes an item from the cart.
	RemoveCart(payload *entity.CartPayload)

	// RemoveAllCartByUser removes all items from the cart for a user.
	// This should only be called when the user is checking out.
	RemoveAllCartByUser(userId string)
}
