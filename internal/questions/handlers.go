package questions

import (
	"core/configs"
	"core/internal/blocks"
	"core/internal/users"

	"net/http"
	"strconv"

	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/quessapp/toolkit/responses"
)

// CreateQuestionHandler is a handler to create a question.
func CreateQuestionHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository, usersRepository *users.UsersRepository, blocksRepository *blocks.BlocksRepository) error {
	payload := CreateQuestionDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	payload.SentBy = authenticatedUserID

	if err := CreateQuestion(handlerCtx, &payload, authenticatedUserID, questionsRepository, usersRepository, blocksRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// GetAllQuestionsHandler is a handler to find all paginated questions.
func GetAllQuestionsHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, questionsRepository *QuestionsRepository) error {
	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	p, err := strconv.Atoi(handlerCtx.C.Query("page"))

	page := int64(p)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	sort := handlerCtx.C.Query("sort")
	filter := handlerCtx.C.Query("filter")

	questions, err := GetAllQuestions(handlerCtx, &page, &sort, &filter, authenticatedUserID, usersRepository, questionsRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, questions)
}

// FindQuestionByIDHandler is a handler to find a question by its id.
func FindQuestionByIDHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository, questionsRepository *QuestionsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	question, err := FindQuestionByID(handlerCtx, ID, authenticatedUserID, questionsRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, question)
}

// DeleteQuestionHandler is a handler to delete a question by its id.
func DeleteQuestionHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	if err := DeleteQuestion(handlerCtx, ID, authenticatedUserID, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}

// HideQuestionHandler is a handler to hide question by its id.
func HideQuestionHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	if err := HideQuestion(handlerCtx, ID, authenticatedUserID, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}

// ReplyQuestionHandler is a handler to reply question by its id.
func ReplyQuestionHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository) error {
	payload := ReplyQuestionDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	payload.ID = ID

	if err := ReplyQuestion(handlerCtx, &payload, authenticatedUserID, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// EditReplyQuestionHandler is a handler to edit reply question by its id.
func EditReplyQuestionHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository) error {
	payload := EditQuestionReplyDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	payload.ID = ID

	if err := EditQuestionReply(handlerCtx, &payload, authenticatedUserID, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}

// RemoveQuestionReplyHandler is a handler to remove reply from question.
func RemoveQuestionReplyHandler(handlerCtx *configs.HandlersCtx, questionsRepository *QuestionsRepository) error {
	ID, err := toolkitEntities.ParseID(handlerCtx.C.Params("id"))

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	if err := RemoveQuestionReply(handlerCtx, ID, authenticatedUserID, questionsRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}
