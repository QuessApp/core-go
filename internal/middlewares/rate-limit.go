package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/kuriozapp/toolkit/responses"
)

// ApplyRateLimitMiddleware applies rate limiter middleware for all routes.
func ApplyRateLimitMiddleware(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s_%s", c.IP(), c.Path())
		},
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return responses.ParseUnsuccesfull(c, fiber.StatusTooManyRequests, "Você fez muitas requisições em pouco tempo. Por favor, aguarde alguns segundos.")
		},
	}))
}
