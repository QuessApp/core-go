package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	appErrors "core/pkg/errors"
	"errors"
)

// CreateQuestion reads payload from request body then try to create a new question in database.
func CreateQuestion(payload entities.Question, questionsRepository *repositories.Questions, usersRepository *repositories.Users) (*entities.Question, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	userToSendQuestion := usersRepository.FindUserByNick(payload.SendTo)

	// TODO: VALIDATE USER IS BLOCKED BY RECEIVER, IS SENDING TO YOURSELF, DID BLOCK RECEIVER, etc.

	if userToSendQuestion.Nick == "" {
		return nil, errors.New(appErrors.USER_NOT_FOUND)
	}

	err := questionsRepository.Create(payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
