package middlewares

import (
	"core/internal/configs"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func shouldProtectRoute(currentRoute string) bool {
	nonProtectedRoutes := []string{"/auth/signin", "/auth/signup"}

	for _, route := range nonProtectedRoutes {
		if route == currentRoute {
			return false
		}
	}

	return true
}

// LoadJWTMiddleware applies JWT middleware for specifics routes.
func LoadJWTMiddleware(app *fiber.App, cfg *configs.Conf) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(cfg.JWTSecret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.ParseUnsuccesfull(c, http.StatusForbidden, err.Error())
		},
		Filter: func(c *fiber.Ctx) bool {
			shouldSkipJWTMiddleware := !shouldProtectRoute(c.OriginalURL())

			return shouldSkipJWTMiddleware
		},
	}))
}
