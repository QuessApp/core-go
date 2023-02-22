package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	validations "core/internal/validations/services"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
func CreateQuestion(payload *entities.Question, questionsRepository *repositories.Questions, blocksRepository *repositories.Blocks, usersRepository *repositories.Users) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	if err := validations.ValidateDidBlockedReceiver(blocksRepository.IsUserBlocked(payload.SendTo)); err != nil {
		return err
	}

	if err := validations.ValidateIsBlockedByReceiver(blocksRepository.IsUserBlocked(payload.SentBy)); err != nil {
		return err
	}

	userToSendQuestion, err := usersRepository.FindUserByID(payload.SendTo)

	if err != nil {
		return err
	}

	if err := validations.ValidateUserExists(userToSendQuestion); err != nil {
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

	if err := questionsRepository.Create(payload); err != nil {
		return err
	}

	return nil
}
