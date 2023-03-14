package users

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all users routes of app.
// It create a group of routes called "users" and assign it the handlers and paths.
func LoadRoutes(AppCtx *configs.AppCtx, usersRepository *UsersRepository) {
	g := AppCtx.App.Group("/users", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	g.Get("/", func(c *fiber.Ctx) error {
		return SearchUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
	g.Get("/me", func(c *fiber.Ctx) error {
		return GetAuthenticatedUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
	g.Put("/me", func(c *fiber.Ctx) error {
		return UpdateUserProfileHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
	g.Patch("/me/avatar", func(c *fiber.Ctx) error {
		return UpdateUserAvatarHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
	g.Get("/:nick", func(c *fiber.Ctx) error {
		return FindUserByNickHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
}
