package carts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewCartRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.CartUseCase,
) {
	cartHandler := NewCartHandler(useCase)

	app.Post("/carts/:id", jwtMiddleware.GuardJWT, cartHandler.AddCart)
	app.Get("/carts", jwtMiddleware.GuardJWT, cartHandler.GetAllCarts)
	app.Delete("/carts/:id", jwtMiddleware.GuardJWT, cartHandler.DeleteCart)
	app.Delete("/carts", jwtMiddleware.GuardJWT, cartHandler.DeleteAlLCart)
}
