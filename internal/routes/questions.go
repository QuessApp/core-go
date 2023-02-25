package routes

import (
	"core/internal/handlers"
	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadQuestionsRoute loads all questions routes of app.
func LoadQuestionsRoute(ctx *AppCtx) {
	g := ctx.App.Group("/questions", middlewares.JWTMiddleware(ctx.App, ctx.Cfg))

	g.Get("/:id", func(c *fiber.Ctx) error {
		return handlers.FindQuestionByIDHandler(c, ctx.QuestionsRepository, ctx.UsersRepository)
	})
	g.Get("", func(c *fiber.Ctx) error {
		return handlers.GetAllQuestionsHandler(c, ctx.QuestionsRepository, ctx.UsersRepository)
	})
	g.Post("", func(c *fiber.Ctx) error {
		return handlers.CreateQuestionHandler(c, ctx.QuestionsRepository, ctx.BlocksRepository, ctx.UsersRepository)
	})
	g.Patch("/hide/:id", func(c *fiber.Ctx) error {
		return handlers.HideQuestionHandler(c, ctx.UsersRepository, ctx.QuestionsRepository)
	})
	g.Delete("/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteQuestionHandler(c, ctx.QuestionsRepository)
	})
}
