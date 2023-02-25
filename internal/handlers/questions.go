package handlers

import (
	"core/internal/dtos"
	"core/internal/repositories"
	"core/internal/services"
	pkg "core/pkg/entities"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"
	"strconv"

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

// GetAllQuestionsHandler is a handler to find all paginated questions.
func GetAllQuestionsHandler(c *fiber.Ctx, questionsRepository *repositories.Questions, usersRepository *repositories.Users) error {
	authenticatedUserId := jwt.GetUserByToken(c).ID

	p, err := strconv.Atoi(c.Query("page"))

	page := int64(p)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	sort := c.Query("sort")
	filter := c.Query("filter")

	questions, err := services.GetAllQuestions(&page, &sort, &filter, authenticatedUserId, questionsRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusOK, questions)
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

// HideQuestionHandler is a handler to hide question by its id.
func HideQuestionHandler(c *fiber.Ctx, usersRepository *repositories.Users, questionsRepository *repositories.Questions) error {
	id, err := pkg.ParseID(c.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(c).ID

	if err := services.HideQuestion(id, authenticatedUserId, questionsRepository, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(c, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(c, http.StatusOK, nil)
}
