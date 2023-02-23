package handlers

import (
	"core/internal/dtos"
	"core/internal/repositories"
	"core/internal/services"
	pkg "core/pkg/entities"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// BlockUserHandler is a handler to block an user.
func BlockUserHandler(c *fiber.Ctx, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) error {
	payload := dtos.BlockUserDTO{}
	id, err := pkg.ParseID(c.Params("id"))

	if err != nil {

		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	payload.BlockedBy = jwt.GetUserByToken(c).ID
	payload.UserToBlock = id

	if err := services.BlockUser(&payload, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, nil)
}
