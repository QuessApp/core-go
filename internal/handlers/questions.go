package handlers

import (
	"core/internal/entities"
	"core/internal/repositories"
	"core/internal/services"
	"core/pkg/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(c *fiber.Ctx, questionsRepository *repositories.Questions, usersRepository *repositories.Users) error {
	payload := entities.Question{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	if err := services.CreateQuestion(payload, questionsRepository, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, nil)
}
