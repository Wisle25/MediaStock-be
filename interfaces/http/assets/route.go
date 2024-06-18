package assets

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewAssetRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.AssetUseCase,
) {
	assetHandler := NewAssetHandler(useCase)

	app.Post("/assets", jwtMiddleware.GuardJWT, assetHandler.AddAsset)
	app.Get("/assets", assetHandler.GetAll)
	app.Get("/assets/:id", assetHandler.GetDetail)
	app.Get("/assets-download/:id", jwtMiddleware.GuardJWT, assetHandler.DownloadAsset)
	app.Put("/assets/:id", jwtMiddleware.GuardJWT, assetHandler.Update)
	app.Delete("/assets/:id", jwtMiddleware.GuardJWT, assetHandler.Delete)
}
