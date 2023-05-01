package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/pkg/i18n"
	"github.com/quessapp/toolkit/responses"
)

const (
	MAX_REQUESTS             = 100
	TIME_TO_RESET_IN_SECONDS = 30
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
			handlerCtx := configs.HandlersCtx{C: c}

			return responses.ParseUnsuccesfull(c, fiber.StatusTooManyRequests, i18n.Translate(&handlerCtx, "max_rate_limit"))
		},
	}))
}
