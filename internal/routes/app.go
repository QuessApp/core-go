package routes

import (
	"core/internal/configs"
	"core/internal/middlewares"
	"core/internal/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadRoutes loads all routes of app.
func LoadRoutes(db *mongo.Database, cfg *configs.Conf, questionsRepository *repositories.Questions, authRepository *repositories.Auth, usersRepository *repositories.Users) {
	app := fiber.New()
	middlewares.LoadMiddlewares(app, cfg)

	LoadAuthRoutes(app, db, authRepository, usersRepository)
	LoadQuestionsRoute(app, db, questionsRepository, usersRepository)

	log.Fatal(app.Listen(cfg.ServerPort))
}
