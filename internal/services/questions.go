package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	validations "core/internal/validations/services"
	pkg "core/pkg/entities"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
func CreateQuestion(payload *entities.Question, authenticatedUserId pkg.ID, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	if err := validations.ValidateIsSendingQuestionToYourself(payload.SendTo.(pkg.ID), authenticatedUserId); err != nil {
		return err
	}

	if err := validations.ValidateDidBlockedReceiver(blocksRepository.IsUserBlocked(payload.SendTo.(pkg.ID))); err != nil {
		return err
	}

	if err := validations.ValidateIsBlockedByReceiver(blocksRepository.IsUserBlocked(payload.SentBy.(pkg.ID))); err != nil {
		return err
	}

	userToSendQuestion, err := usersRepository.FindUserByID(payload.SendTo.(pkg.ID))

	if err != nil {
		return err
	}

	if err := validations.ValidateUserExists(userToSendQuestion); err != nil {
		return err
	}

	userThatIsSendingQuestion, err := usersRepository.FindUserByID(payload.SentBy.(pkg.ID))

	if err != nil {
		return err
	}

	if err := usersRepository.DecrementLimit(payload.SentBy.(pkg.ID)); err != nil {
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

	return nil
}

// FindQuestionByID finds for a question in database by question id.
func FindQuestionByID(id pkg.ID, authenticatedUserId pkg.ID, questionsRepository *repositories.Questions, usersRepository *repositories.Users) (*entities.Question, error) {
	foundQuestion := questionsRepository.FindByID(id)

	if err := validations.ValidateQuestionExists(foundQuestion); err != nil {
		return nil, err
	}

	if err := validations.ValidateQuestionIsSentForMe(foundQuestion, authenticatedUserId); err != nil {
		return nil, err
	}

	questionOwner, err := usersRepository.FindUserByID(foundQuestion.SentBy.(pkg.ID))

	if err != nil {
		return nil, err
	}

	foundQuestion.SentBy = questionOwner.GetBasicInfos()

	return foundQuestion.MapAnonymousFields(), nil
}
