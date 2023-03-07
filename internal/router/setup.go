package router

import (
	"core/configs"
	"core/docs"
	"core/internal/auth"
	"core/internal/blocks"
	"core/internal/middlewares"
	"core/internal/questions"
	"core/internal/settings"
	"core/internal/users"
	"log"
)

// Setup setups middlewares, initializes routes, etc.
func Setup(AppCtx *configs.AppCtx, authRepository *auth.AuthRepository, usersRepository *users.UsersRepository, blocksRepository *blocks.BlocksRepository, questionsRepository *questions.QuestionsRepository) {
	middlewares.ApplyMiddlewares(AppCtx.App, AppCtx.Cfg)

	auth.LoadRoutes(AppCtx, authRepository, usersRepository)
	questions.LoadRoutes(AppCtx, usersRepository, questionsRepository, blocksRepository)
	blocks.LoadRoutes(AppCtx, usersRepository, blocksRepository)
	docs.LoadRoutes(AppCtx)
	users.LoadRoutes(AppCtx, usersRepository)
	settings.LoadRoutes(AppCtx, usersRepository)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
