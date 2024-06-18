package favorites

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type FavoriteHandler struct {
	useCase *use_case.FavoriteUseCase
}

func NewFavoriteHandler(useCase *use_case.FavoriteUseCase) *FavoriteHandler {
	return &FavoriteHandler{useCase}
}

func (h *FavoriteHandler) AddFavorite(c *fiber.Ctx) error {
	// Payload
	var payload entity.FavoritePayload
	payload.AssetId = c.Params("id")
	payload.UserId = c.Locals("userInfo").(entity.User).Id

	// Use Case
	h.useCase.ExecuteAdd(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully added favorite",
	})
}

func (h *FavoriteHandler) GetAllFavorites(c *fiber.Ctx) error {
	// Payload
	userId := c.Locals("userInfo").(entity.User).Id

	// Use Case
	favorites := h.useCase.ExecuteGetAll(userId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   favorites,
	})
}

func (h *FavoriteHandler) DeleteFavorite(c *fiber.Ctx) error {
	// Payload
	var payload entity.FavoritePayload
	payload.AssetId = c.Params("id")
	payload.UserId = c.Locals("userInfo").(entity.User).Id

	// Use Case
	h.useCase.ExecuteRemove(&payload)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully deleted favorite",
	})
}
