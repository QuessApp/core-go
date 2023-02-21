package routes

import (
	"core/internal/handlers"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadQuestionsRoute(app *fiber.App, db *mongo.Database, questionsRepository *repositories.Questions, usersRepository *repositories.Users) {
	g := app.Group("/questions")

	g.Post("", func(c *fiber.Ctx) error {
		return handlers.CreateQuestionHandler(c, questionsRepository, usersRepository)
	})
}
