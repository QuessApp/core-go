package settings

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/middlewares"
	"github.com/quessapp/core-go/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes is responsible for setting up the routes related to settings in the Fiber app.
// AppCtx is the application context.
// usersRepository is the repository for users.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *users.UsersRepository) {
	g := AppCtx.App.Group("/settings", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Patch("/preferences", func(c *fiber.Ctx) error {
		return UpdatePreferencesHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
}
