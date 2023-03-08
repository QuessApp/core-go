package users

import (
	"core/configs"

	"core/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all users routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *UsersRepository) {
	g := AppCtx.App.Group("/users", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Get("/", func(c *fiber.Ctx) error {
		return SearchUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
	g.Get("/me", func(c *fiber.Ctx) error {
		return GetAuthenticatedUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
}
