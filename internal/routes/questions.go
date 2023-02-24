package routes

import (
	"core/internal/configs"
	"core/internal/handlers"
	"core/internal/middlewares"
	"core/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadQuestionsRoute loads all questions routes of app.
func LoadQuestionsRoute(app *fiber.App, db *mongo.Database, cfg *configs.Conf, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) {
	g := app.Group("/questions", middlewares.JWTMiddleware(app, cfg))

	g.Post("", func(c *fiber.Ctx) error {
		h := handlers.CreateQuestionHandler(c, questionsRepository, blocksRepository, usersRepository)

		return h
	})
	g.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.FindQuestionByIDHandler(c, questionsRepository, usersRepository)
	})
}
