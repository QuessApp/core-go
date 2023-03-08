package questions

import (
	"core/configs"
	"core/internal/blocks"
	"core/internal/middlewares"

	"core/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all questions routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *users.UsersRepository, questionsRepository *QuestionsRepository, blocksRepository *blocks.BlocksRepository) {
	g := AppCtx.App.Group("/questions", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Get("/:id", func(c *fiber.Ctx) error {
		return FindQuestionByIDHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, questionsRepository)
	})
	g.Get("/", func(c *fiber.Ctx) error {
		return GetAllQuestionsHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository, questionsRepository)
	})
	g.Post("/", func(c *fiber.Ctx) error {
		return CreateQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository, usersRepository, blocksRepository)
	})
	g.Patch("/hide/:id", func(c *fiber.Ctx) error {
		return HideQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	})
	g.Delete("/:id", func(c *fiber.Ctx) error {
		return DeleteQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	})
	g.Patch("/reply/:id", func(c *fiber.Ctx) error {
		return ReplyQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	})
	g.Delete("/reply/:id", func(c *fiber.Ctx) error {
		return RemoveQuestionReplyHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	})
	g.Patch("/reply/edit/:id", func(c *fiber.Ctx) error {
		return EditReplyQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	})
}
