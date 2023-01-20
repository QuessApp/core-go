package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Fiber does not handle panics by default. To recover from a panic thrown by any handler in the stack, you need to include the Recover middleware below.
// See https://docs.gofiber.io/guide/error-handling
func Recover(app *fiber.App) {
	app.Use(recover.New())
}
