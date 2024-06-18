package ratings

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type RatingHandler struct {
	useCase *use_case.RatingUseCase
}

func NewRatingHandler(useCase *use_case.RatingUseCase) *RatingHandler {
	return &RatingHandler{useCase}
}

func (h *RatingHandler) AddRating(c *fiber.Ctx) error {
	// Payload
	var payload entity.CreateRatingPayload
	_ = c.BodyParser(&payload)

	payload.UserId = c.Locals("userInfo").(entity.User).Id

	// Use Case
	id := h.useCase.ExecuteCreate(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully added rating",
		"data":    id,
	})
}

func (h *RatingHandler) GetRatingsByAsset(c *fiber.Ctx) error {
	// Payload
	assetId := c.Params("assetId")

	// Use Case
	ratings := h.useCase.ExecuteGetByAsset(assetId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   ratings,
	})
}

func (h *RatingHandler) DeleteRating(c *fiber.Ctx) error {
	// Payload
	ratingId := c.Params("id")

	// Use Case
	h.useCase.ExecuteDelete(ratingId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully deleted rating",
	})
}
