package api

import (
	"core/cmd/api/middlewares"
	"core/configs"
	"core/internal/auth"
	"core/internal/blocks"
	"core/internal/questions"
	"core/internal/users"
	"log"
)

// LoadRoutes loads all app routes.
func LoadRoutes(AppCtx *configs.AppCtx, authRepository *auth.AuthRepository, usersRepository *users.UsersRepository, blocksRepository *blocks.BlocksRepository, questionsRepository *questions.QuestionsRepository) {
	middlewares.LoadMiddlewares(AppCtx.App, AppCtx.Cfg)
	auth.LoadRoutes(AppCtx, authRepository, usersRepository)
	questions.LoadRoutes(AppCtx, usersRepository, questionsRepository, blocksRepository)
	blocks.LoadRoutes(AppCtx, usersRepository, blocksRepository)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
