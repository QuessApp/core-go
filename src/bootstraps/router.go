package bootstraps

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func initRouterConfig(router *fiber.App) {
	InitRoutes(router)
	InitMiddlewares(router)
}

// InitRouter starts a new http server.
func InitRouter() {
	PORT := os.Getenv("PORT")

	router := fiber.New()
	initRouterConfig(router)

	log.Printf("App running in port %s \n", PORT)

	log.Fatal(router.Listen(PORT))
}
