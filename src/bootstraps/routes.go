package bootstraps

import (
	"core/src/routes"

	"github.com/gofiber/fiber/v2"
)

// InitRoutes inits all routes in our app.
func InitRoutes(router *fiber.App) {
	routes.App(router)
}
