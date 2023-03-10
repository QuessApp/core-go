package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/quessapp/toolkit/responses"
)

const (
	MAX_REQUESTS             = 100
	TIME_TO_RESET_IN_SECONDS = 30
	ERROR_MESSAGE            = "Você fez muitas requisições em pouco tempo. Por favor, aguarde alguns segundos."
)

// ApplyRateLimitMiddleware applies rate limiter middleware for all routes.
func ApplyRateLimitMiddleware(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max: MAX_REQUESTS,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s_%s", c.IP(), c.Path())
		},
		Expiration: TIME_TO_RESET_IN_SECONDS * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return responses.ParseUnsuccesfull(c, fiber.StatusTooManyRequests, ERROR_MESSAGE)
		},
	}))
}
