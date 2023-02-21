package routes

import (
	"core/internal/handlers"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadAuthRoutes(app *fiber.App, db *mongo.Database, authRepository *repositories.Auth, usersRepository *repositories.Users) {
	g := app.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.SignUpUserHandler(c, usersRepository, authRepository)
	})
}
