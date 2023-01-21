package controllers

import (
	"core/src/database"
	helpers "core/src/helpers/responses"
	usecases "core/src/usecases/auth"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser handles the request to register an user.
func RegisterUser(c *fiber.Ctx) error {
	db, _ := database.Connect()

	createdUser, err := usecases.RegisterUser(c, db)

	return helpers.ParseControllerResponse(c, err, createdUser)
}
