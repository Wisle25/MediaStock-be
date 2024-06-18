package transactions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewTransactionRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.TransactionUseCase,
) {
	transactionHandler := NewTransactionHandler(useCase)

	app.Post("/transactions", jwtMiddleware.GuardJWT, transactionHandler.AddTransaction)
	app.Get("/transactions", jwtMiddleware.GuardJWT, transactionHandler.GetTransactionByUser)
	app.Get("/transactions/:id", jwtMiddleware.GuardJWT, transactionHandler.GetTransactionById)
}
