package routes

import (
	"core/internal/configs"
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadQuestionsRoute loads all questions routes of app.
func LoadQuestionsRoute(AppCtx *configs.AppCtx) {
	g := AppCtx.App.Group("/questions", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.FindQuestionByIDHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Get("", func(c *fiber.Ctx) error {
		return handlers.GetAllQuestionsHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Post("", func(c *fiber.Ctx) error {
		return handlers.CreateQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Patch("/hide/:id", func(c *fiber.Ctx) error {
		return handlers.HideQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Delete("/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Patch("/reply/:id", func(c *fiber.Ctx) error {
		return handlers.ReplyQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
}
