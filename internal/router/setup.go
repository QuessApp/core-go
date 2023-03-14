package router

import (
	"log"

	"github.com/quessapp/core-go/docs"

	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/internal/middlewares"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/internal/settings"
	"github.com/quessapp/core-go/internal/users"
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
