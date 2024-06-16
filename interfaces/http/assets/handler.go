﻿package assets

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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

	// Use Case
	assets := h.ExecuteGetAll(listCount, pageList)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   assets,
	})
}

func (h *AssetHandler) GetDetail(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")

	// Use Case
	asset := h.ExecuteGetDetail(id)

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
	userId := c.Locals("user_id").(string)
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
	userId := c.Locals("user_id").(string)

	// Use Case
	h.VerifyAccess(userId, id)
	h.ExecuteDelete(id)

	// Response
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully delete asset!",
	})
}
