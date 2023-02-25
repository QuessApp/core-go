package routes

import (
	"core/cmd/app/entities"
	"core/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// LoadAuthRoutes loads all auth routes of app.
func LoadAuthRoutes(AppCtx *entities.AppCtx) {
	g := AppCtx.App.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignUpUserHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return handlers.SignInUserHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
}
