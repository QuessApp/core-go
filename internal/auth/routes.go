package auth

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"

	"github.com/gofiber/fiber/v2"
)

// LoadRoutes is a function that sets up the routes for the auth API.
// It takes in an AppCtx, a AuthRepository, and a UserRepository.
func LoadRoutes(AppCtx *configs.AppCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) {
	g := AppCtx.App.Group("/auth")

	g.Post("/signup", func(c *fiber.Ctx) error {
		return SignUpUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Post("/signin", func(c *fiber.Ctx) error {
		return SignInUserHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Post("/refresh", func(c *fiber.Ctx) error {
		return RefreshTokenHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Delete("/logout", func(c *fiber.Ctx) error {
		return LogoutHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Post("/forgot-password", func(c *fiber.Ctx) error {
		return ForgotPasswordHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
	g.Put("/reset-password", func(c *fiber.Ctx) error {
		return ResetPasswordHandler(&configs.HandlersCtx{C: c, AppCtx: *AppCtx}, authRepository, usersRepository)
	})
}
