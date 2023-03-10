package questions

import (
	"core/configs"
	"core/internal/blocks"
	"core/internal/emails"

	"core/internal/users"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// CreateQuestion creates a new question in the system and sends an email notification to the recipient if enabled.
// It returns an error if any validation checks fail or if there is an issue with creating the question.
func CreateQuestion(handlerCtx *configs.HandlersCtx, payload *CreateQuestionDTO, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository, usersRepository *users.UsersRepository, blocksRepository *blocks.BlocksRepository) error {
	if err := IsInvalidSendToID(payload); err != nil {
		return err
	}

	if err := payload.Validate(); err != nil {
		return err
	}

	if err := IsSendingQuestionToYourself(payload.SendTo, authenticatedUserId); err != nil {
		return err
	}

	if err := blocks.DidBlockedReceiver(blocksRepository.IsUserBlocked(payload.SendTo)); err != nil {
		return err
	}

	payload.SentBy = authenticatedUserId

	if err := blocks.IsBlockedByReceiver(blocksRepository.IsUserBlocked(payload.SentBy)); err != nil {
		return err
	}

	userToSendQuestion := usersRepository.FindUserByID(payload.SendTo)

	if err := users.UserExists(userToSendQuestion); err != nil {
		return err
	}

	userThatIsSendingQuestion := usersRepository.FindUserByID(payload.SentBy)

	if err := ReachedPostsLimitToCreateQuestion(userThatIsSendingQuestion); err != nil {
		return err
	}

	if err := users.DecrementUserLimit(userThatIsSendingQuestion.ID, usersRepository); err != nil {
		return err
	}

	// TODO: add flag in database
	if userThatIsSendingQuestion.IsShadowBanned {
		// fake question, dont create
		// record in database
		return nil
	}

	if err := questionsRepository.Create(payload); err != nil {
		return err
	}

	if userToSendQuestion.EnableAPPEmails {
		emails.SendEmailNewQuestionReceived(handlerCtx.AppCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.SendEmailsQueue, payload.Content, payload.IsAnonymous, userToSendQuestion, userThatIsSendingQuestion)
	}

	if err := users.UpdateLastPublishedAt(userThatIsSendingQuestion, usersRepository); err != nil {
		return err
	}

	if err := users.ResetLimit(userThatIsSendingQuestion, usersRepository); err != nil {
		return err
	}

	return nil
}

// FindQuestionByID retrieves a question with the provided ID from the questions repository and returns
// a Question object and an error. Before returning the question, it is checked if the question exists
// and if the authenticated user has permission to view the question.
//
// If the question is owned by an anonymous user, the function maps the anonymous fields of the question
// and returns the mapped question.
func FindQuestionByID(handlerCtx *configs.HandlersCtx, id, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository, usersRepository *users.UsersRepository) (*Question, error) {
	q := questionsRepository.FindQuestionByID(id)

	if err := QuestionExists(q); err != nil {
		return nil, err
	}

	if err := QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return nil, err
	}

	questionOwner := usersRepository.FindUserByID(q.SentBy.(toolkitEntities.ID))

	u := users.User{
		ID:        questionOwner.ID,
		Name:      questionOwner.Name,
		Nick:      questionOwner.Nick,
		AvatarURL: questionOwner.AvatarURL,
	}

	q.SentBy = u

	return q.MapAnonymousFields(), nil
}

// GetAllQuestions retrieves a paginated list of questions from the questions repository and returns
// a PaginatedQuestions object and an error. The pagination, sorting and filtering parameters are optional
// and have default values. The authenticated user ID is used to determine which questions the user has
// permission to view. If there are no questions that match the filter, an empty array is returned.
//
// For each question, the function checks if the question is owned by an anonymous user. If so, it sets
// the "SentBy" field to nil. Otherwise, it retrieves the user who owns the question from the users
// repository and maps the user fields to a new User object, which is assigned to the "SentBy" field of
// the question. The function then returns a PaginatedQuestions object that contains the list of questions
// and the total count of questions that match the filter.
func GetAllQuestions(handlerCtx *configs.HandlersCtx, page *int64, sort, filter *string, authenticatedUserId toolkitEntities.ID, usersRepository *users.UsersRepository, questionsRepository *QuestionsRepository) (*PaginatedQuestions, error) {
	if *page == 0 {
		*page = 1
	}

	if *sort == "" {
		*sort = "asc"
	}

	if *filter == "" {
		*filter = "all"
	}

	questions, err := questionsRepository.GetAll(page, sort, filter, authenticatedUserId)

	if err != nil {
		return nil, err
	}

	if len(*questions.Questions) == 0 {
		return &PaginatedQuestions{
			Questions:  &[]Question{},
			TotalCount: 0,
		}, nil
	}

	var allQuestions []Question

	for _, q := range *questions.Questions {
		if q.IsAnonymous {
			q.SentBy = nil
		} else {
			u := usersRepository.FindUserByID(q.SentBy.(toolkitEntities.ID))

			q.SentBy = users.User{
				ID:        u.ID,
				Nick:      u.Nick,
				Name:      u.Name,
				AvatarURL: u.AvatarURL,
			}
		}

		allQuestions = append(allQuestions, q)
	}

	result := PaginatedQuestions{
		Questions:  &allQuestions,
		TotalCount: questions.TotalCount,
	}

	return &result, err
}

// DeleteQuestion retrieves the question with the provided ID from the questions repository and checks if it exists.
// If the question exists, the function checks if the authenticated user has permission to delete the question.
// If the user has permission, the function deletes the question from the repository.
// If the question does not exist or the user does not have permission to delete the question, the function returns an error.
func DeleteQuestion(handlerCtx *configs.HandlersCtx, id, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
	foundQuestion := questionsRepository.FindQuestionByID(id)

	if err := QuestionExists(foundQuestion); err != nil {
		return err
	}

	if err := CanUserDeleteQuestion(foundQuestion, authenticatedUserId); err != nil {
		return err
	}

	if err := questionsRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

// HideQuestion is a function that takes in a handler context, question id, authenticated user id, and a questions repository as arguments.
// It retrieves the question from the questions repository using the id, checks if the question exists, and if it can be hidden by the authenticated user.
// It also checks if the authenticated user can view the question and if the question has not been previously hidden by the receiver.
// If all checks pass, it calls the questions repository's Hide function to hide the question.
func HideQuestion(handlerCtx *configs.HandlersCtx, id, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
	q := questionsRepository.FindQuestionByID(id)

	if err := QuestionExists(q); err != nil {
		return err
	}

	if err := CanHideQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := IsHiddenByReceiver(q.IsHiddenByReceiver); err != nil {
		return err
	}

	if err := questionsRepository.Hide(id); err != nil {
		return err
	}

	return nil
}

// ReplyQuestion is a function that takes in a handler context, a reply question DTO, authenticated user id, and a questions repository as arguments.
// It validates the reply question DTO, retrieves the question from the questions repository using the id, and checks if the question can be viewed by the authenticated user.
// It also checks if the question has not already been replied to and if the authenticated user can reply to the question.
// If all checks pass, it calls the questions repository's Reply function to add the reply to the question.
func ReplyQuestion(handlerCtx *configs.HandlersCtx, payload *ReplyQuestionDTO, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	q := questionsRepository.FindQuestionByID(payload.ID)

	if err := QuestionExists(q); err != nil {
		return err
	}

	if err := QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := IsAlreadyReplied(q); err != nil {
		return err
	}

	if err := CanReply(q, authenticatedUserId); err != nil {
		return err
	}

	if err := questionsRepository.Reply(payload); err != nil {
		return err
	}

	return nil
}

// EditQuestionReply is a function that takes in a handler context, an edit question reply DTO, authenticated user id, and a questions repository as arguments.
// It validates the edit question reply DTO, retrieves the question from the questions repository using the id, and checks if the authenticated user can reply to the question.
// It also checks if the question has already been replied to, if the authenticated user has not reached the limit for editing the reply, and if the question is not yet replied.
// If all checks pass, it sets the old content and creation date of the question in the DTO and calls the questions repository's EditReply function to edit the reply.
func EditQuestionReply(handlerCtx *configs.HandlersCtx, payload *EditQuestionReplyDTO, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	q := questionsRepository.FindQuestionByID(payload.ID)

	if err := QuestionExists(q); err != nil {
		return err
	}

	if err := QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := CanReply(q, authenticatedUserId); err != nil {
		return err
	}

	if err := ReachedLimitToEditReply(q); err != nil {
		return err
	}

	if err := IsQuestionNotRepliedYet(q); err != nil {
		return err
	}

	payload.OldContent = q.Content
	payload.OldContentCreatedAt = q.RepliedAt

	if err := questionsRepository.EditReply(payload); err != nil {
		return err
	}

	return nil
}

// RemoveQuestionReply is a function that takes in a handler context, a question id, authenticated user id, and a questions repository as arguments.
// It retrieves the question from the questions repository using the id, and checks if the authenticated user can view the question and if the question has been replied to.
// If all checks pass, it calls the questions repository's RemoveReply function to remove the reply
func RemoveQuestionReply(handlerCtx *configs.HandlersCtx, id, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
	q := questionsRepository.FindQuestionByID(id)

	if err := QuestionExists(q); err != nil {
		return err
	}

	if err := QuestionCanViewQuestion(q, authenticatedUserId); err != nil {
		return err
	}

	if err := IsQuestionNotRepliedYet(q); err != nil {
		return err
	}

	if err := questionsRepository.RemoveReply(id); err != nil {
		return err
	}

	return nil
}