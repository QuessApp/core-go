package bootstraps

import (
	"log"

	"github.com/joho/godotenv"
)

// InitEnv loads ``.env`` file.
func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
