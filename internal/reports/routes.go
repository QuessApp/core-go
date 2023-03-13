package reports

import (
	"core/configs"
	"core/internal/middlewares"
	"core/internal/questions"
	"core/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all questions routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository, reportsRepository *ReportsRepository) {
	g := AppCtx.App.Group("/reports", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/", func(c *fiber.Ctx) error {
		return CreateReportHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository, usersRepository, reportsRepository)
	})

	// g.Delete("/:id", func(c *fiber.Ctx) error {
	// 	return DeleteQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	// })
}
