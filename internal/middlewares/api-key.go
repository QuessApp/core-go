package middlewares

import (
	"core/configs"
	"log"
	"net/http"

	"github.com/quessapp/toolkit/middlewares"
	"github.com/quessapp/toolkit/responses"

	"github.com/gofiber/fiber/v2"
)

// ApplyAPIKeyMiddleware applies API key middleware for all routes.
func ApplyAPIKeyMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(middlewares.New(middlewares.Config{
		Next: func(c *fiber.Ctx) bool {
			isDev := cfg.Env == "development"

			if isDev {
				log.Println("[DEV] For development purposes like debugging the API key middleware is disabled.")
			}

			return isDev
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.ParseUnsuccesfull(c, http.StatusForbidden, err.Error())
		},
		Key: cfg.APIKey,
	}))
}
