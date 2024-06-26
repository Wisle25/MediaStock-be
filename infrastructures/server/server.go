package server

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/infrastructures/cache"
	"github.com/wisle25/media-stock-be/infrastructures/container"
	"github.com/wisle25/media-stock-be/infrastructures/emails"
	"github.com/wisle25/media-stock-be/infrastructures/file_statics"
	"github.com/wisle25/media-stock-be/infrastructures/generator"
	"github.com/wisle25/media-stock-be/infrastructures/services"
	"github.com/wisle25/media-stock-be/interfaces/http/assets"
	"github.com/wisle25/media-stock-be/interfaces/http/carts"
	"github.com/wisle25/media-stock-be/interfaces/http/comments"
	"github.com/wisle25/media-stock-be/interfaces/http/favorites"
	"github.com/wisle25/media-stock-be/interfaces/http/middlewares"
	"github.com/wisle25/media-stock-be/interfaces/http/ratings"
	"github.com/wisle25/media-stock-be/interfaces/http/transactions"
	"github.com/wisle25/media-stock-be/interfaces/http/users"
)

func errorHandling(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	status := "error"
	code := fiber.StatusInternalServerError
	message := err.Error()

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		status = "fail"
		code = e.Code
		message = e.Message
	}

	// Send custom error
	return c.Status(code).JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}

func CreateServer(config *commons.Config) *fiber.App {
	// Load Services
	db := services.ConnectDB(config)
	redis := services.ConnectRedis(config)
	minio, bucketName := services.NewMinio(config)

	// Server
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandling,
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowMethods:     "POST,PUT,GET,DELETE,OPTIONS,PATCH",
		AllowCredentials: true,
	}))

	// Global Dependencies
	redisCache := cache.NewRedisCache(redis)
	uuidGenerator := generator.NewUUIDGenerator()
	validation := services.NewValidation()
	minioFileUpload := file_statics.NewMinioFileUpload(minio, uuidGenerator, bucketName)
	vipsFileProcessing := file_statics.NewVipsFileProcessing()
	emailService := emails.NewStmpEmailService(config)

	// Use Cases
	userUseCase := container.NewUserContainer(
		config,
		db,
		redisCache,
		uuidGenerator,
		vipsFileProcessing,
		minioFileUpload,
		emailService,
		validation,
	)
	assetUseCase := container.NewAssetContainer(uuidGenerator, db, vipsFileProcessing, minioFileUpload, validation, config)
	favoriteUseCase := container.NewFavoriteContainer(uuidGenerator, db)
	cartUseCase := container.NewCartContainer(uuidGenerator, db)
	transactionUseCase := container.NewTransactionContainer(uuidGenerator, db)
	ratingUseCase := container.NewRatingContainer(uuidGenerator, db, validation)
	commentUseCase := container.NewCommentContainer(uuidGenerator, db, validation)

	// Custom Middleware
	jwtMiddleware := middlewares.NewJwtMiddleware(userUseCase)

	// Router
	users.NewUserRouter(app, jwtMiddleware, userUseCase, config)
	assets.NewAssetRouter(app, jwtMiddleware, assetUseCase)
	favorites.NewFavoriteRouter(app, jwtMiddleware, favoriteUseCase)
	carts.NewCartRouter(app, jwtMiddleware, cartUseCase)
	transactions.NewTransactionRouter(app, jwtMiddleware, transactionUseCase)
	ratings.NewRatingRouter(app, jwtMiddleware, ratingUseCase)
	comments.NewCommentRouter(app, jwtMiddleware, commentUseCase)

	return app
}
