package assets

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
	"strconv"
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

	payload.OwnerId = c.Locals("userInfo").(entity.User).Id
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

func (h *AssetHandler) GetAll(c *fiber.Ctx) error {
	// Payload
	listCount, err := strconv.Atoi(c.Query("listCount"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	pageList, err := strconv.Atoi(c.Query("pageList"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	defaultId, _ := uuid.NewV7() // This is to prevent invalid UUID by SQL
	userId := c.Query("userId", defaultId.String())

	// Use Case
	assets := h.ExecuteGetAll(listCount, pageList, userId)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   assets,
	})
}

func (h *AssetHandler) GetDetail(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")
	defaultId, _ := uuid.NewV7() // This is to prevent invalid UUID by SQL
	userId := c.Query("userId", defaultId.String())

	// Use Case
	asset := h.ExecuteGetDetail(id, userId)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   asset,
	})
}

func (h *AssetHandler) Update(c *fiber.Ctx) error {
	var err error

	// Payload
	id := c.Params("id")
	userId := c.Locals("userInfo").(entity.User).Id
	var payload entity.AddAssetPayload
	_ = c.BodyParser(&payload)

	payload.File, err = c.FormFile("asset")
	if err != nil {
		return fmt.Errorf("upload asset: %v", err)
	}

	// Use Case
	h.VerifyAccess(userId, id)
	h.ExecuteUpdate(id, &payload)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully update asset!",
	})
}

func (h *AssetHandler) Delete(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")
	userId := c.Locals("userInfo").(entity.User).Id

	// Use Case
	h.VerifyAccess(userId, id)
	h.ExecuteDelete(id)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully delete asset!",
	})
}

func (h *AssetHandler) DownloadAsset(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")
	userId := c.Locals("userInfo").(entity.User).Id

	// Use Case
	title, fileBuffer := h.AssetUseCase.ExecuteDownload(id, userId)

	// Response
	c.Set("Content-Disposition", "attachment; filename="+title)
	c.Set("Content-Type", "application/octet-stream")

	// Send the file to the client
	return c.Send(fileBuffer)
}
