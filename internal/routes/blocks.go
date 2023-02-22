package routes

import (
	"core/internal/configs"
	"core/internal/handlers"
	"core/internal/middlewares"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadBlocksRoutes loads all blocks routes of app.
func LoadBlocksRoutes(app *fiber.App, db *mongo.Database, cfg *configs.Conf, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) {
	g := app.Group("/blocks", middlewares.JWTMiddleware(app, cfg))

	g.Post("/user/:id", func(c *fiber.Ctx) error {
		return handlers.BlockUserHandler(c, usersRepository, blocksRepository)
	})
}
