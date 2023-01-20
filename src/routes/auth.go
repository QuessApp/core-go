package routes

import (
	"core/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(router *fiber.App) {
	auth := router.Group("/auth")

	auth.Post("/signup", controllers.RegisterUser)
}
