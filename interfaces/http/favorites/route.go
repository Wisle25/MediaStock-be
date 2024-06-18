package favorites

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewFavoriteRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.FavoriteUseCase,
) {
	favoriteHandler := NewFavoriteHandler(useCase)

	app.Post("/favorites/:id", jwtMiddleware.GuardJWT, favoriteHandler.AddFavorite)
	app.Get("/favorites", jwtMiddleware.GuardJWT, favoriteHandler.GetAllFavorites)
	app.Delete("/favorites/:id", jwtMiddleware.GuardJWT, favoriteHandler.DeleteFavorite)
}
