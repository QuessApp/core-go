package handlers

import (
	"core/internal/configs"
	"core/internal/dtos"
	"core/internal/repositories"
	"core/internal/services"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(c *fiber.Ctx, cfg *configs.Conf, usersRepository *repositories.Users, authRepository *repositories.Auth) error {
	payload := dtos.SignUpUserDTO{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignUp(cfg, &payload, usersRepository, authRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, u)
}

func SignInUserHandler(c *fiber.Ctx, cfg *configs.Conf, usersRepository *repositories.Users) error {
	payload := dtos.SignInUserDTO{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignIn(cfg, payload.Nick, payload.Password, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, u)
}
