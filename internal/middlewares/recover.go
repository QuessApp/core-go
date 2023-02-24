package middlewares

import (
	"core/internal/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// LoadRecoverMiddleware applies recover middleware for all routes.
// See https://docs.gofiber.io/api/middleware/recover/
func LoadRecoverMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			shouldSkip := cfg.Env == "development"

			if shouldSkip {
				log.Println("[DEV] For development purpouses like debugging the recover middleware is disabled.")
			}

			return shouldSkip
		},
	}))
}
