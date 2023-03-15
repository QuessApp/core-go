package reports

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/middlewares"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes is responsible for setting up the routes related to reports in the Fiber app.
// AppCtx is the application context.
// questionsRepository is the repository for questions.
// usersRepository is the repository for users.
// reportsRepository is the repository for reports.
func LoadRoutes(AppCtx *configs.AppCtx, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository, reportsRepository *ReportsRepository) {
	g := AppCtx.App.Group("/reports", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Post("/", func(c *fiber.Ctx) error {
		return CreateReportHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository, usersRepository, reportsRepository)
	})

	// g.Delete("/:id", func(c *fiber.Ctx) error {
	// 	return DeleteQuestionHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, questionsRepository)
	// })
}
