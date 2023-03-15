package blocks

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/middlewares"
	"github.com/quessapp/core-go/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes is a function that sets up the routes for the blocks API.
// It takes in an AppCtx, a UsersRepository, and a BlocksRepository.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) {
	g := AppCtx.App.Group("/blocks", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return BlockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, blocksRepository)
	})
	g.Patch("/user/:id", func(c *fiber.Ctx) error {
		return UnblockUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, blocksRepository)
	})
}
