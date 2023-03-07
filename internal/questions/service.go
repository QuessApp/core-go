package questions

import (
	"core/configs"
	"core/internal/blocks"

	"core/internal/users"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
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

	// TODO: maybe use Go routines?
	if err := users.DecrementUserLimit(userThatIsSendingQuestion.ID, usersRepository); err != nil {
		return err
	}

	if userThatIsSendingQuestion.IsShadowBanned {
		// fake question, dont create
		// record in database
		return nil
	}

	if err := questionsRepository.Create(payload); err != nil {
		return err
	}

	// TODO: update user lastPublishAt field.
	if userToSendQuestion.EnableAppEmails {
		go SendEmailNewQuestionReceived(handlerCtx.AppCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.SendEmailsQueue, payload, userToSendQuestion, userThatIsSendingQuestion)
	}

	return nil
}

// FindQuestionByID finds for a question in database by question id.
func FindQuestionByID(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository, usersRepository *users.UsersRepository) (*Question, error) {
	foundQuestion := questionsRepository.FindQuestionByID(id)

	if err := QuestionExists(foundQuestion); err != nil {
		return nil, err
	}

	if err := QuestionCanViewQuestion(foundQuestion, authenticatedUserId); err != nil {
		return nil, err
	}

	questionOwner := usersRepository.FindUserByID(foundQuestion.SentBy.(toolkitEntities.ID))

	u := users.User{
		ID:        questionOwner.ID,
		Name:      questionOwner.Name,
		Nick:      questionOwner.Nick,
		AvatarURL: questionOwner.AvatarURL,
	}

	foundQuestion.SentBy = u

	return foundQuestion.MapAnonymousFields(), nil
}

// GetAllQuestions gets all paginated questions from database.
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

// DeleteQuestion deletes a question from database by id.
func DeleteQuestion(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
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

// HideQuestion hides a question.
func HideQuestion(handlerCtx *configs.HandlersCtx, id toolkitEntities.ID, authenticatedUserId toolkitEntities.ID, questionsRepository *QuestionsRepository) error {
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

// ReplyQuestion replies a question.
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

// EditQuestionReply edits a question reply.
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
