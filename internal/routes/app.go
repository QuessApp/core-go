package routes

import (
	"core/internal/configs"
	"core/internal/middlewares"
	"core/internal/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// AppCtx is a global model for app. It defines the router, db, config, repositories, etc.
// Use AppCtx to avoid long function params.
type AppCtx struct {
	App                 *fiber.App
	DB                  *mongo.Database
	Cfg                 *configs.Conf
	QuestionsRepository *repositories.Questions
	BlocksRepository    *repositories.Blocks
	UsersRepository     *repositories.Users
	AuthRepository      *repositories.Auth
}

// LoadRoutes loads all routes of app.
func LoadRoutes(AppCtx *AppCtx) {
	middlewares.LoadMiddlewares(AppCtx.App, AppCtx.Cfg)

	LoadAuthRoutes(AppCtx)
	LoadQuestionsRoute(AppCtx)
	LoadBlocksRoutes(AppCtx)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
