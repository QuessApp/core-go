package middlewares

import (
	"core/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// ApplyRecoverMiddleware applies recover middleware for all routes.
// See https://docs.gofiber.io/api/middleware/recover/
func ApplyRecoverMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			isDev := cfg.Env == "development"

			if isDev {
				log.Println("[DEV] For development purposes like debugging the recover middleware is disabled.")
			}

			return isDev
		},
	}))
}
