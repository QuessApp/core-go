package handlers

import (
	"core/internal/entities"
	"core/internal/repositories"
	"core/internal/services"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(c *fiber.Ctx, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) error {
	payload := entities.Question{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	payload.SentBy = jwt.GetUserByToken(c).ID

	if err := services.CreateQuestion(&payload, questionsRepository, blocksRepository, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, nil)
}
