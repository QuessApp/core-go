package middlewares

import (
	"core/internal/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// LoadCORSMiddleware applies CORS middleware for all routes.
func LoadCORSMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
	}))
}
