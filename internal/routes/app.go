package routes

import (
	"core/internal/configs"
	"core/internal/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(db *mongo.Database, cfg *configs.Conf, authRepository *repositories.Auth, usersRepository *repositories.Users) {
	app := fiber.New()

	LoadAuthRoutes(app, db, authRepository, usersRepository)

	log.Fatal(app.Listen(cfg.ServerPort))
}
