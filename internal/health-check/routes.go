package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/users"
)

// LoadRoutes is a function to load health check routes.
// This function will take four parameters, an AppCtx, an AuthRepository, a QuestionsRepository, and a UsersRepository.
func LoadRoutes(AppCtx *configs.AppCtx, authRepository *auth.AuthRepository, questionsRepository *questions.QuestionsRepository, usersRepository *users.UsersRepository) {
	g := AppCtx.App.Group("/health-check")

	g.Get("/", func(c *fiber.Ctx) error {
		return RunHealthCheckHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, questionsRepository, usersRepository)
	})
}
