package middlewares

import (
	"github.com/quessapp/core-go/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// ApplyCORSMiddleware applies CORS middleware for all routes.
func ApplyCORSMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
	}))
}
