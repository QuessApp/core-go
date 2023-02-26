package middlewares

import (
	"core/configs"

	"github.com/gofiber/fiber/v2"
)

// LoadMiddlewares applies middlewares to fiber router.
func LoadMiddlewares(app *fiber.App, cfg *configs.Conf) {
	LoadRecoverMiddleware(app, cfg)
	LoadAPIKeyMiddleware(app, cfg)
	LoadCORSMiddleware(app, cfg)
	LoadLoggerMiddleware(app)
}
