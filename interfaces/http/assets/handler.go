package assets

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type AssetHandler struct {
	*use_case.AssetUseCase
}

func NewAssetHandler(useCase *use_case.AssetUseCase) *AssetHandler {
	return &AssetHandler{
		useCase,
	}
}

func (h *AssetHandler) AddAsset(c *fiber.Ctx) error {
	var err error

	// Payload
	var payload entity.AddAssetPayload
	_ = c.BodyParser(&payload)

	payload.OwnerId = c.Locals("user_id").(string)
	payload.File, err = c.FormFile("asset")
	if err != nil {
		return fmt.Errorf("upload asset: %v", err)
	}
	
	// Use Case
	returnedId := h.ExecuteAdd(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    returnedId,
		"message": "Successfully create asset!",
	})
}
