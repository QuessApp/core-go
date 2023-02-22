package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// LoadRecoverMiddleware applies recover middleware for all routes.
// See https://docs.gofiber.io/api/middleware/recover/
func LoadRecoverMiddleware(app *fiber.App) {
	app.Use(recover.New())
}
