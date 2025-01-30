package server

import (
	"company-microservice/api/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitServer() *fiber.App {
	// --------------------Start
	app := fiber.New()

	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())
	app.Use(logger.New())

	router.InitRouter(app)

	app.Listen(":" + os.Getenv("PORT"))

	return app
}
