package handlers

import (
	"core/internal/entities"
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
	payload := entities.BlockedUser{}
	id := c.Query("id")

	payload.BlockedBy = jwt.GetUserByToken(c).ID
	payload.UserToBlock, _ = pkg.ParseID(id)

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	if err := services.BlockUser(&payload, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, nil)
}
