package questions

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/blocks"
	"github.com/quessapp/core-go/internal/users"

	"net/http"
	"strconv"

	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/quessapp/toolkit/responses"
)

// CreateQuestionHandler creates a new question using the provided payload.
// It takes three parameters, a HandlerCtx, a QuestionsRepository, a UsersRepository, and a BlocksRepository.
// It returns an error if the creation is unsuccessful.
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

// GetAllQuestionsHandler retrieves all questions based on the provided filters and returns them as a paginated list.
// It takes three parameters, a HandlerCtx, a UsersRepository, and a QuestionsRepository.
// It returns an error if the retrieval is unsuccessful.
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

// FindQuestionByIDHandler retrieves a question based on the provided ID and returns it.
// It takes three parameters, a HandlerCtx, a UsersRepository, and a QuestionsRepository.
// It returns an error if the retrieval is unsuccessful.
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

// DeleteQuestionHandler handles the request to delete a question with the given ID.
// It requires a HandlersCtx object and a QuestionsRepository object as input parameters.
// It returns an error if the ID cannot be parsed or if the question cannot be deleted.
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

// HideQuestionHandler handles the request to hide a question with the given ID.
// It requires a HandlersCtx object and a QuestionsRepository object as input parameters.
// It returns an error if the ID cannot be parsed or if the question cannot be hidden.
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

// ReplyQuestionHandler handles the request to reply to a question with the given ID.
// It requires a HandlersCtx object and a QuestionsRepository object as input parameters.
// It returns an error if the request payload cannot be parsed, if the ID cannot be parsed, or if the question cannot be replied to.
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

// EditReplyQuestionHandler handles the request to edit a reply to a question with the given ID.
// It requires a HandlersCtx object and a QuestionsRepository object as input parameters.
// It returns an error if the request payload cannot be parsed, if the ID cannot be parsed, or if the reply cannot be edited.
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

// RemoveQuestionReplyHandler handles the request to remove a reply to a question with the given ID.
// It requires a HandlersCtx object and a QuestionsRepository object as input parameters.
// It returns an error if the ID cannot be parsed or if the reply cannot be removed.
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
