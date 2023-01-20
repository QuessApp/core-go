package routes

import (
	"core/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// App will handle defauilt/generics app routes like ``/``, ``/ping``, etc...
func App(router *fiber.App) {
	router.Post("/ping", controllers.PingAppController)
}
