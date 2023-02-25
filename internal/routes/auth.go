package routes

import (
	"core/internal/configs"
	"core/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// LoadAuthRoutes loads all auth routes of app.
func LoadAuthRoutes(AppCtx *configs.AppCtx) {
	g := AppCtx.App.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignUpUserHandler(&configs.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return handlers.SignInUserHandler(&configs.HandlersContext{C: c, AppCtx: *AppCtx})
	})
}
