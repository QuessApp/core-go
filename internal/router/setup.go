package router

import (
	"core/configs"
	"core/docs"
	"core/internal/auth"
	"core/internal/blocks"
	"core/internal/middlewares"
	"core/internal/questions"
	"core/internal/reports"
	"core/internal/settings"
	"core/internal/users"
	"log"
)

// TODO: may we can attach all repositories into AppCtx to avoid long params?

// Setup setups middlewares, initializes routes, etc.
func Setup(AppCtx *configs.AppCtx, authRepository *auth.AuthRepository, usersRepository *users.UsersRepository, blocksRepository *blocks.BlocksRepository, questionsRepository *questions.QuestionsRepository, reportsRepository *reports.ReportsRepository) {
	middlewares.ApplyMiddlewares(AppCtx.App, AppCtx.Cfg)

	auth.LoadRoutes(AppCtx, authRepository, usersRepository)
	questions.LoadRoutes(AppCtx, usersRepository, questionsRepository, blocksRepository)
	blocks.LoadRoutes(AppCtx, usersRepository, blocksRepository)
	users.LoadRoutes(AppCtx, usersRepository)
	settings.LoadRoutes(AppCtx, usersRepository)
	reports.LoadRoutes(AppCtx, questionsRepository, usersRepository, reportsRepository)
	docs.LoadRoutes(AppCtx)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
