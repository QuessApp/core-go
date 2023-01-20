package bootstraps

import (
	"core/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

// InitMiddlewares inits app middlewares
func InitMiddlewares(router *fiber.App) {
	middlewares.Recover(router)
}
