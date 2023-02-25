package routes

import (
	"core/cmd/app/entities"
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadBlocksRoutes loads all blocks routes of app.
func LoadBlocksRoutes(AppCtx *entities.AppCtx) {
	g := AppCtx.App.Group("/blocks", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return handlers.BlockUserHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Patch("/user/:id", func(c *fiber.Ctx) error {
		return handlers.UnblockUserHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
}
