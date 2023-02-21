package middlewares

import (
	"core/internal/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// LoadMiddlewares applies middlewares to fiber router.
func LoadMiddlewares(app *fiber.App, cfg *configs.Conf) {
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
	}))
}
