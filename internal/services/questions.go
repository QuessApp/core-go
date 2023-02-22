package services

import (
	"core/internal/dtos"
	"core/internal/entities"
	"core/internal/repositories"
	validations "core/internal/validations/services"
	pkg "core/pkg/entities"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
func CreateQuestion(payload *dtos.CreateQuestionDTO, authenticatedUserId pkg.ID, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) error {

	if err := payload.Validate(); err != nil {
		return err
	}

	if err := validations.IsSendingQuestionToYourself(payload.SendTo, authenticatedUserId); err != nil {
		return err
	}

	if err := validations.DidBlockedReceiver(blocksRepository.IsUserBlocked(payload.SendTo)); err != nil {
		return err
	}

	payload.SentBy = authenticatedUserId

	if err := validations.IsBlockedByReceiver(blocksRepository.IsUserBlocked(payload.SentBy)); err != nil {
		return err
	}

	userToSendQuestion, err := usersRepository.FindUserByID(payload.SendTo)

	if err != nil {
		return err
	}

	if err := validations.UserExists(userToSendQuestion); err != nil {
		return err
	}

	userThatIsSendingQuestion, err := usersRepository.FindUserByID(payload.SentBy)

	if err != nil {
		return err
	}

	if err := usersRepository.DecrementLimit(payload.SentBy); err != nil {
		return err
	}

	if userThatIsSendingQuestion.IsShadowBanned {
		// fake question, dont create
		// record in database
		return nil
	}

	if err := validations.ReachedPostsLimitToCreateQuestion(userThatIsSendingQuestion); err != nil {
		return err
	}

	if err := questionsRepository.Create(payload); err != nil {
		return err
	}

	return nil
}

// FindQuestionByID finds for a question in database by question id.
func FindQuestionByID(id pkg.ID, authenticatedUserId pkg.ID, questionsRepository *repositories.Questions, usersRepository *repositories.Users) (*entities.Question, error) {
	foundQuestion := questionsRepository.FindByID(id)

	if err := validations.QuestionExists(foundQuestion); err != nil {
		return nil, err
	}

	if err := validations.QuestionIsSentForMe(foundQuestion, authenticatedUserId); err != nil {
		return nil, err
	}

	questionOwner, err := usersRepository.FindUserByID(foundQuestion.SentBy.(pkg.ID))

	if err != nil {
		return nil, err
	}

	u := entities.User{
		ID:        questionOwner.ID,
		Name:      questionOwner.Name,
		Nick:      questionOwner.Nick,
		AvatarURL: questionOwner.AvatarURL,
	}

	foundQuestion.SentBy = u

	return foundQuestion.MapAnonymousFields(), nil
}
