package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ApplyLoggerMiddleware applies a logger middleware for all routes.
func ApplyLoggerMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
}
