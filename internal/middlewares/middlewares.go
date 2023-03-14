package middlewares

import (
	"github.com/quessapp/core-go/configs"

	"github.com/gofiber/fiber/v2"
)

// ApplyMiddlewares applies middlewares to fiber router.
func ApplyMiddlewares(app *fiber.App, cfg *configs.Conf) {
	ApplyAPIKeyMiddleware(app, cfg)
	ApplyCORSMiddleware(app, cfg)
	ApplyLoggerMiddleware(app)
	ApplyRecoverMiddleware(app, cfg)
	ApplyRateLimitMiddleware(app)
}
