package blocks

import (
	"core/cmd/api/middlewares"
	"core/configs"
	"core/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all blocks routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) {
	g := AppCtx.App.Group("/blocks", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return BlockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, blocksRepository)
	})
	g.Patch("/user/:id", func(c *fiber.Ctx) error {
		return UnblockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, blocksRepository)
	})
}
