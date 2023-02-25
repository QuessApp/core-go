package routes

import (
	"core/cmd/app/entities"
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadQuestionsRoute loads all questions routes of app.
func LoadQuestionsRoute(AppCtx *entities.AppCtx) {
	g := AppCtx.App.Group("/questions", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.FindQuestionByIDHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Get("", func(c *fiber.Ctx) error {
		return handlers.GetAllQuestionsHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Post("", func(c *fiber.Ctx) error {
		return handlers.CreateQuestionHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Patch("/hide/:id", func(c *fiber.Ctx) error {
		return handlers.HideQuestionHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
	g.Delete("/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteQuestionHandler(&entities.HandlersContext{C: c, AppCtx: *AppCtx})
	})
}
