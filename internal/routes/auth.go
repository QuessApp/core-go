package routes

import (
	"core/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// LoadAuthRoutes loads all auth routes of app.
func LoadAuthRoutes(AppCtx *AppCtx) {
	g := AppCtx.App.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignUpUserHandler(c, AppCtx.Cfg, AppCtx.UsersRepository, AppCtx.AuthRepository)
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return handlers.SignInUserHandler(c, AppCtx.Cfg, AppCtx.UsersRepository)
	})
}
