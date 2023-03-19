package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

// ApplyHelmetMiddleware applies helmet middleware for all routes.
func ApplyHelmetMiddleware(app *fiber.App) {
	app.Use(helmet.New())
}
