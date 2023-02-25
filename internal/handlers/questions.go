package handlers

import (
	"core/cmd/app/entities"
	"core/internal/dtos"
	"core/internal/services"
	pkg "core/pkg/entities"
	"core/pkg/jwt"
	"core/pkg/responses"
	"net/http"
	"strconv"
)

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(handlerCtx *entities.HandlersContext) error {
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
func GetAllQuestionsHandler(handlerCtx *entities.HandlersContext) error {
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
func FindQuestionByIDHandler(handlerCtx *entities.HandlersContext) error {
	id, err := pkg.ParseID(handlerCtx.C.Params("id"))

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
func DeleteQuestionHandler(handlerCtx *entities.HandlersContext) error {
	id, err := pkg.ParseID(handlerCtx.C.Params("id"))

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
func HideQuestionHandler(handlerCtx *entities.HandlersContext) error {
	id, err := pkg.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := jwt.GetUserByToken(handlerCtx.C).ID

	if err := services.HideQuestion(id, authenticatedUserId, handlerCtx.QuestionsRepository, handlerCtx.UsersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}
