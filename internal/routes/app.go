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
func LoadRoutes(db *mongo.Database, cfg *configs.Conf, questionsRepository *repositories.Questions, authRepository *repositories.Auth, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) {
	app := fiber.New()
	middlewares.LoadMiddlewares(app, cfg)

	LoadAuthRoutes(app, db, cfg, authRepository, usersRepository)
	LoadQuestionsRoute(app, db, cfg, questionsRepository, blocksRepository, usersRepository)
	LoadBlocksRoutes(app, db, cfg, usersRepository, blocksRepository)

	log.Fatal(app.Listen(cfg.ServerPort))
}
