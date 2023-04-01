package middlewares

import (
	"net/http"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/toolkit/responses"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// JWTMiddleware applies JWT middleware for specifics routes.
func JWTMiddleware(app *fiber.App, cfg *configs.Conf) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(cfg.JWT.Secret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return responses.ParseUnsuccesfull(c, http.StatusForbidden, err.Error())
		},
	})
}
