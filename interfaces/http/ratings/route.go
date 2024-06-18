package ratings

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewRatingRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.RatingUseCase,
) {
	ratingHandler := NewRatingHandler(useCase)

	app.Post("/ratings", jwtMiddleware.GuardJWT, ratingHandler.AddRating)
	app.Get("/ratings/:assetId", ratingHandler.GetRatingsByAsset)
	app.Delete("/ratings/:id", jwtMiddleware.GuardJWT, ratingHandler.DeleteRating)
}
