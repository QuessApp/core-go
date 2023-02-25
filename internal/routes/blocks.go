package routes

import (
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadBlocksRoutes loads all blocks routes of app.
func LoadBlocksRoutes(AppCtx *AppCtx) {
	g := AppCtx.App.Group("/blocks", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return handlers.BlockUserHandler(c, AppCtx.UsersRepository, AppCtx.BlocksRepository)
	})
	g.Patch("/user/:id", func(c *fiber.Ctx) error {
		return handlers.UnblockUserHandler(c, AppCtx.UsersRepository, AppCtx.BlocksRepository)
	})
}
