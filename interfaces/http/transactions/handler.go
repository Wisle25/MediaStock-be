package transactions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/domains/entity"
)

type TransactionHandler struct {
	useCase *use_case.TransactionUseCase
}

func NewTransactionHandler(useCase *use_case.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		useCase,
	}
}

func (h *TransactionHandler) AddTransaction(c *fiber.Ctx) error {
	// Payload (Receiving total amount and assets id)
	var payload entity.CreateTransactionPayload
	_ = c.BodyParser(&payload)

	payload.UserID = c.Locals("userInfo").(entity.UserToken).UserId

	// Use Case
	returnedId := h.useCase.ExecuteCreate(&payload)

	// Response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"data":    returnedId,
		"message": "Successfully make transaction!",
	})
}

func (h *TransactionHandler) GetTransactionByUser(c *fiber.Ctx) error {
	// Payload
	userId := c.Locals("userInfo").(entity.UserToken).UserId

	// Use Case
	transactions := h.useCase.ExecuteGetByUser(userId)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   transactions,
	})
}

func (h *TransactionHandler) GetTransactionById(c *fiber.Ctx) error {
	// Payload
	id := c.Params("id")

	// Use Case
	transaction := h.useCase.ExecuteGetByID(id)

	// Response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   transaction,
	})
}
