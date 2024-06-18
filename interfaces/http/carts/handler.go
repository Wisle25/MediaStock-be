package carts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type CartHandler struct {
	useCase *use_case.CartUseCase
}

func NewCartHandler(useCase *use_case.CartUseCase) *CartHandler {
	return &CartHandler{useCase}
}

func (h *CartHandler) AddCart(c *fiber.Ctx) error {
	// Payload
	var payload entity.CartPayload
	payload.AssetId = c.Params("id")
	payload.UserId = c.Locals("userInfo").(entity.UserToken).UserId

	// Use Case
	h.useCase.ExecuteAdd(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully added cart",
	})
}

func (h *CartHandler) GetAllCarts(c *fiber.Ctx) error {
	// Payload
	userId := c.Locals("userInfo").(entity.UserToken).UserId

	// Use Case
	carts := h.useCase.ExecuteGetAll(userId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   carts,
	})
}

func (h *CartHandler) DeleteCart(c *fiber.Ctx) error {
	// Payload
	var payload entity.CartPayload
	payload.AssetId = c.Params("id")
	payload.UserId = c.Locals("userInfo").(entity.UserToken).UserId

	// Use Case
	h.useCase.ExecuteRemove(&payload)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully delete a cart",
	})
}
