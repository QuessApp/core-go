package routes

import (
	"core/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Questions(router *fiber.App) {
	questions := router.Group("/questions")

	questions.Post("/", controllers.CreateQuestion)
}
