package middlewares

import (
	"core/internal/configs"
	"log"
	"net/http"

	"github.com/kuriozapp/toolkit/middlewares"
	"github.com/kuriozapp/toolkit/responses"

	"github.com/gofiber/fiber/v2"
)

// LoadAPIKeyMiddleware applies API key middleware for all routes.
func LoadAPIKeyMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(middlewares.New(middlewares.Config{
		Next: func(c *fiber.Ctx) bool {
			isDev := cfg.Env == "development"

			if isDev {
				log.Println("[DEV] For development purpouses like debugging the API key middleware is disabled.")
			}

			return isDev
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.ParseUnsuccesfull(c, http.StatusForbidden, err.Error())
		},
		Key: cfg.APIKey,
	}))
}
