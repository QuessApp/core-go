package routes

import (
	"core/internal/configs"
	"core/internal/handlers"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadAuthRoutes loads all auth routes of app.
func LoadAuthRoutes(app *fiber.App, db *mongo.Database, cfg *configs.Conf, authRepository *repositories.Auth, usersRepository *repositories.Users) {
	g := app.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignUpUserHandler(c, cfg, usersRepository, authRepository)
	})

	g.Post("/signin", func(c *fiber.Ctx) error {
		return handlers.SignInUserHandler(c, cfg, usersRepository)
	})
}
