package handlers

import (
	"core/internal/configs"
	"core/internal/entities"
	"core/internal/repositories"
	"core/internal/services"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(c *fiber.Ctx, cfg *configs.Conf, usersRepository *repositories.Users, authRepository *repositories.Auth) error {
	payload := entities.User{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignUp(&payload, usersRepository, authRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	accessToken, err := jwt.CreateAccessToken(u, cfg)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	refreshToken, err := jwt.CreateRefreshToken(u, cfg)

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
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return responses.ParseSuccessful(c, http.StatusCreated, data)
}
