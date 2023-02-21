package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	appErrors "core/pkg/errors"
	"errors"
)

func CreateQuestion(payload entities.Question, usersRepository *repositories.Users) (*entities.Question, error) {
	userToSendQuestion := usersRepository.FindUserByNick(payload.SendTo.Nick)

	// TODO: VALIDATE USER IS BLOCKED BY RECEIVER, IS SENDING TO YOURSELF, DID BLOCK RECEIVER, etc.

	if userToSendQuestion.Nick == "" {
		return nil, errors.New(appErrors.USER_NOT_FOUND)
	}

	return nil, nil
}
