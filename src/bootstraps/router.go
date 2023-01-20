package bootstraps

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// InitRouter starts a new http server.
func InitRouter() {
	PORT := os.Getenv("PORT")

	router := fiber.New()

	InitRoutes(router)

	log.Printf("App running in port %s \n", PORT)

	log.Fatal(router.Listen(PORT))
}
