package routes

import (
	"core/internal/configs"
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadBlocksRoutes loads all blocks routes of app.
func LoadBlocksRoutes(AppCtx *configs.AppCtx) {
	g := AppCtx.App.Group("/blocks", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return handlers.BlockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Patch("/user/:id", func(c *fiber.Ctx) error {
		return handlers.UnblockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
}
