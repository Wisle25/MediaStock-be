package comments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewCommentRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.CommentUseCase,
) {
	commentHandler := NewCommentHandler(useCase)

	app.Post("/comments", jwtMiddleware.GuardJWT, commentHandler.AddComment)
	app.Get("/comments/:assetId", commentHandler.GetCommentsByAsset)
	app.Put("/comments/:id", jwtMiddleware.GuardJWT, commentHandler.UpdateComment)
	app.Delete("/comments/:id", jwtMiddleware.GuardJWT, commentHandler.DeleteComment)
}
