package routes

import (
	"core/internal/configs"
	"core/internal/middlewares"
	"log"
)

// LoadRoutes loads all routes of app.
func LoadRoutes(AppCtx *configs.AppCtx) {
	middlewares.LoadMiddlewares(AppCtx.App, AppCtx.Cfg)

	LoadAuthRoutes(AppCtx)
	LoadQuestionsRoute(AppCtx)
	LoadBlocksRoutes(AppCtx)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
