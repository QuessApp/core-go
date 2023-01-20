package controllers

import (
	"core/src/database"
	"core/src/exceptions"
	"core/src/helpers"
	"core/src/models"
	"core/src/usecases"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser handles the request to register an user.
func RegisterUser(c *fiber.Ctx) error {
	db, _ := database.Connect()

	createdUser, err := usecases.RegisterUser(c, db)

	if err != nil {
		return exceptions.NewHttpException(c, models.Response{
			Message: fmt.Sprint(err),
			Status:  400,
		})
	}

	return helpers.ParseSuccessResponse(c, models.Response{
		Data:   createdUser,
		Status: 201,
	})
}
