package handlers

import (
	"core/internal/entities"
	"core/internal/repositories"
	"core/internal/services"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(c *fiber.Ctx, usersRepository *repositories.Users, authRepository *repositories.Auth) error {
	payload := entities.User{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignUp(&payload, usersRepository, authRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	data := &entities.ResponseWithUser{
		User: &entities.User{
			ID:        u.ID,
			AvatarURL: u.AvatarURL,
			Name:      u.Name,
			Email:     u.Email,
		},
	}

	return responses.ParseSuccessful(c, http.StatusCreated, data)
}
