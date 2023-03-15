package auth

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all auth routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) {
	g := AppCtx.App.Group("/auth")

	g.Get("/", func(c *fiber.Ctx) error {
		return RedirectToAuthPageHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})

	g.Get("/callback", func(c *fiber.Ctx) error {
		return AuthenticateHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx})
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return SignInUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
}
