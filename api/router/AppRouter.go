package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	CompanyRouter(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("ok")
	})

	app.All("/*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON("not found")
	})
}
