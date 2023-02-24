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

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(c *fiber.Ctx, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) error {
	payload := dtos.CreateQuestionDTO{}

	if err := c.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(c).ID

	payload.SentBy = authenticatedUserId

	if err := services.CreateQuestion(&payload, authenticatedUserId, questionsRepository, blocksRepository, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusCreated, nil)
}

// FindQuestionByIDHandler is a handler to find a question by its id.
func FindQuestionByIDHandler(c *fiber.Ctx, questionsRepository *repositories.Questions, usersRepository *repositories.Users) error {
	id, err := pkg.ParseID(c.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(c).ID

	question, err := services.FindQuestionByID(id, authenticatedUserId, questionsRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusOK, question)
}

// DeleteQuestionHandler is a handler to delete a question by its id.
func DeleteQuestionHandler(c *fiber.Ctx, questionsRepository *repositories.Questions) error {
	id, err := pkg.ParseID(c.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(c).ID

	if err := services.DeleteQuestion(id, authenticatedUserId, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusOK, nil)
}
