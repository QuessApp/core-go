package handlers

import (
	"core/internal/configs"
	"core/internal/dtos"
	"core/internal/services"
	"core/pkg/jwt"
	"net/http"
	"strconv"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"github.com/kuriozapp/toolkit/responses"
)

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(handlerCtx *configs.HandlersContext) error {
	payload := dtos.CreateQuestionDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	payload.SentBy = authenticatedUserId

	if err := services.CreateQuestion(&payload, authenticatedUserId, handlerCtx.QuestionsRepository, handlerCtx.BlocksRepository, handlerCtx.UsersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// GetAllQuestionsHandler is a handler to find all paginated questions.
func GetAllQuestionsHandler(handlerCtx *configs.HandlersContext) error {
	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	p, err := strconv.Atoi(handlerCtx.C.Query("page"))

	page := int64(p)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	sort := handlerCtx.C.Query("sort")
	filter := handlerCtx.C.Query("filter")

	questions, err := services.GetAllQuestions(&page, &sort, &filter, authenticatedUserId, handlerCtx.QuestionsRepository, handlerCtx.UsersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, questions)
}

// FindQuestionByIDHandler is a handler to find a question by its id.
func FindQuestionByIDHandler(handlerCtx *configs.HandlersContext) error {
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	question, err := services.FindQuestionByID(id, authenticatedUserId, handlerCtx.QuestionsRepository, handlerCtx.UsersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, question)
}

// DeleteQuestionHandler is a handler to delete a question by its id.
func DeleteQuestionHandler(handlerCtx *configs.HandlersContext) error {
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	if err := services.DeleteQuestion(id, authenticatedUserId, handlerCtx.QuestionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}

// HideQuestionHandler is a handler to hide question by its id.
func HideQuestionHandler(handlerCtx *configs.HandlersContext) error {
	id, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	if err := services.HideQuestion(id, authenticatedUserId, handlerCtx.QuestionsRepository, handlerCtx.UsersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}
