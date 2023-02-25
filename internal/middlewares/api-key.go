package middlewares

import (
	"core/internal/configs"
	pkgMiddlewares "core/pkg/middlewares"
	"core/pkg/responses"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// LoadAPIKeyMiddleware applies API key middleware for all routes.
func LoadAPIKeyMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(pkgMiddlewares.New(pkgMiddlewares.Config{
		Next: func(c *fiber.Ctx) bool {
			isDev := cfg.Env == "development"

			if isDev {
				log.Println("[DEV] For development purpouses like debugging the recover API key middleware is disabled.")
			}

			return isDev
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.ParseUnsuccesfull(c, http.StatusForbidden, err.Error())
		},
		Key: cfg.APIKey,
	}))
}
