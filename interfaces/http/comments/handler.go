package comments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type CommentHandler struct {
	useCase *use_case.CommentUseCase
}

func NewCommentHandler(useCase *use_case.CommentUseCase) *CommentHandler {
	return &CommentHandler{useCase}
}

func (h *CommentHandler) AddComment(c *fiber.Ctx) error {
	// Payload
	var payload entity.CreateCommentPayload
	_ = c.BodyParser(&payload)
	payload.UserId = c.Locals("userInfo").(entity.User).Id

	// Use Case
	commentId := h.useCase.ExecuteCreate(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Comment added successfully",
		"data":    commentId,
	})
}

func (h *CommentHandler) GetCommentsByAsset(c *fiber.Ctx) error {
	// Payload
	assetId := c.Params("assetId")

	// Use Case
	comments := h.useCase.ExecuteGetByAsset(assetId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   comments,
	})
}

func (h *CommentHandler) UpdateComment(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")
	var payload entity.EditCommentPayload
	_ = c.BodyParser(&payload)

	// Use Case
	h.useCase.ExecuteUpdate(id, &payload)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Comment updated successfully",
	})
}

func (h *CommentHandler) DeleteComment(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")

	// Use Case
	h.useCase.ExecuteDelete(id)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Comment deleted successfully",
	})
}
