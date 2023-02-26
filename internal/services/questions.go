package services

import (
	"core/internal/configs"
	"core/internal/dtos"
	"core/internal/entities"
	validations "core/internal/validations/services"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
func CreateQuestion(handlerCtx *configs.HandlersCtx, payload *dtos.CreateQuestionDTO, authenticatedUserId toolkitEntities.ID) error {
	if err := validations.IsInvalidSendToID(payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return err
	}

	if err := validations.IsSendingQuestionToYourself(payload.SendTo, authenticatedUserId); err != nil {
		return err
	}

	if err := validations.DidBlockedReceiver(handlerCtx.BlocksRepository.IsUserBlocked(payload.SendTo)); err != nil {
		return err
	}

	payload.SentBy = authenticatedUserId

	if err := validations.IsBlockedByReceiver(handlerCtx.BlocksRepository.IsUserBlocked(payload.SentBy)); err != nil {
		return err
	}

	userToSendQuestion := handlerCtx.UsersRepository.FindUserByID(payload.SendTo)

	if err := validations.UserExists(userToSendQuestion); err != nil {
		return err
	}

	userThatIsSendingQuestion := handlerCtx.UsersRepository.FindUserByID(payload.SentBy)

	if err := validations.ReachedPostsLimitToCreateQuestion(userThatIsSendingQuestion); err != nil {
		return err
	}

	// TODO: maybe use Go routines?
	if err := DecrementUserLimit(userThatIsSendingQuestion.ID, handlerCtx.UsersRepository); err != nil {
		return err
	}

	if userThatIsSendingQuestion.IsShadowBanned {
		// fake question, dont create
		// record in database
		return nil
	}

	if err := handlerCtx.QuestionsRepository.Create(payload); err != nil {
		return err
	}

	// TODO: update user lastPublishAt field.

	if userToSendQuestion.EnableAppEmails {
		go SendNewQuestionReceivedEmail(handlerCtx.AppCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.SendEmailsQueue, payload, userToSendQuestion, userThatIsSendingQuestion)
	}

	return nil
}

// FindQuestionByID finds for a question in database by question id.
func FindQuestionByID(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID) (*entities.Question, error) {
	foundQuestion := handlerCtx.QuestionsRepository.FindByID(id)

	if err := validations.QuestionExists(foundQuestion); err != nil {
		return nil, err
	}

	if err := validations.QuestionCanViewQuestion(foundQuestion, authenticatedUserId); err != nil {
		return nil, err
	}

	questionOwner := handlerCtx.UsersRepository.FindUserByID(foundQuestion.SentBy.(toolkitEntities.ID))

	u := entities.User{
		ID:        questionOwner.ID,
		Name:      questionOwner.Name,
		Nick:      questionOwner.Nick,
		AvatarURL: questionOwner.AvatarURL,
	}

	foundQuestion.SentBy = u

	return foundQuestion.MapAnonymousFields(), nil
}

// GetAllQuestions gets all paginated questions from database.
func GetAllQuestions(handlerCtx *configs.HandlersCtx, page *int64, sort, filter *string, authenticatedUserId toolkitEntities.ID) (*entities.PaginatedQuestions, error) {
	if *page == 0 {
		*page = 1
	}

	if *sort == "" {
		*sort = "asc"
	}

	if *filter == "" {
		*filter = "all"
	}

	questions, err := handlerCtx.QuestionsRepository.GetAll(page, sort, filter, authenticatedUserId)

	if err != nil {
		return nil, err
	}

	var allQuestions []entities.Question

	for _, q := range *questions.Questions {
		if q.IsAnonymous {
			q.SentBy = nil
		} else {
			u := handlerCtx.UsersRepository.FindUserByID(q.SentBy.(toolkitEntities.ID))

			q.SentBy = entities.User{
				ID:        u.ID,
				Nick:      u.Nick,
				Name:      u.Name,
				AvatarURL: u.AvatarURL,
			}
		}

		allQuestions = append(allQuestions, q)
	}

	result := entities.PaginatedQuestions{
		Questions:  &allQuestions,
		TotalCount: questions.TotalCount,
	}

	return &result, err
}

// DeleteQuestion deletes a question from database by id.
func DeleteQuestion(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID) error {
	foundQuestion := handlerCtx.QuestionsRepository.FindByID(id)

	if err := validations.QuestionExists(foundQuestion); err != nil {
		return err
	}

	if err := validations.CanUserDeleteQuestion(foundQuestion, authenticatedUserId); err != nil {
		return err
	}

	if err := handlerCtx.QuestionsRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

// HideQuestion hides a question.
func HideQuestion(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID) error {
	q := handlerCtx.QuestionsRepository.FindByID(id)

	if err := validations.QuestionExists(q); err != nil {
		return err
	}

	if err := validations.CanHideQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := validations.QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := validations.IsHiddenByReceiver(q.IsHiddenByReceiver); err != nil {
		return err
	}

	if err := handlerCtx.QuestionsRepository.Hide(id); err != nil {
		return err
	}

	return nil
}
