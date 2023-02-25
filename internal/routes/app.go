package routes

import (
	"core/cmd/app/entities"
	"core/internal/middlewares"
	"log"
)

// LoadRoutes loads all routes of app.
func LoadRoutes(AppCtx *entities.AppCtx) {
	middlewares.LoadMiddlewares(AppCtx.App, AppCtx.Cfg)

	LoadAuthRoutes(AppCtx)
	LoadQuestionsRoute(AppCtx)
	LoadBlocksRoutes(AppCtx)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.ServerPort))
}
