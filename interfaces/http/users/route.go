package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
)

func NewUserRouter(
	app *fiber.App,
	jwtMiddleware *middlewares.JwtMiddleware,
	useCase *use_case.UserUseCase,
	config *commons.Config,
) {
	userHandler := NewUserHandler(useCase, config)

	app.Post("/users", userHandler.AddUser)
	app.Post("/auths", userHandler.Login)
	app.Get("/auths", jwtMiddleware.GuardJWT, userHandler.GetLoggedUser)
	app.Put("/auths", userHandler.RefreshToken)
	app.Delete("/auths", jwtMiddleware.GuardJWT, userHandler.Logout)
	app.Get("/activate", userHandler.ActivateAccount)
	app.Get("/users/:id", userHandler.GetUserById)
	app.Put("/users/:id", jwtMiddleware.GuardJWT, userHandler.UpdateUserById)
}
