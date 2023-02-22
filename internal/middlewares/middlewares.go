package middlewares

import (
	"core/internal/configs"

	"github.com/gofiber/fiber/v2"
)

// LoadMiddlewares applies middlewares to fiber router.
func LoadMiddlewares(app *fiber.App, cfg *configs.Conf) {
	LoadRecoverMiddleware(app)
	LoadCORSMiddleware(app, cfg)
	LoadLoggerMiddleware(app)
}
