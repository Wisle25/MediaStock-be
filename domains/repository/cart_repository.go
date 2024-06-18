package repository

import "github.com/wisle25/media-stock-be/domains/entity"

type CartRepository interface {
	// AddToCart should raise panic if both asset or user is not existed
	AddToCart(payload *entity.CartPayload)

	// GetAllCarts returns favorited asset by user id
	GetAllCarts(userId string) []entity.Cart

	// RemoveCart remove favorite from user
	RemoveCart(payload *entity.CartPayload)
}
