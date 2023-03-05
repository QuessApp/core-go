package auth

import (
	"core/configs"
	"core/internal/middlewares"
	"core/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes loads all auth routes of app.
//
// It create routes and assign handlers to each route.
func LoadRoutes(AppCtx *configs.AppCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) {
	g := AppCtx.App.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return SignUpUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return SignInUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})

	// TODO: refactor this
	withMiddleware := AppCtx.App.Group("/auth", middlewares.JWTMiddleware(AppCtx.App, AppCtx.Cfg))

	withMiddleware.Get("/me", func(c *fiber.Ctx) error {
		return GetAuthenticatedUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, usersRepository)
	})
}
