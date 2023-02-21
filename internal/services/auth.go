package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	appErrors "core/pkg/errors"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser reads payload from request body then try to register a new user in database.
func RegisterUser(payload entities.User, usersRepository *repositories.Users, authRepository *repositories.Auth) (*entities.User, error) {
	payload.Format()

	err := payload.Validate()

	if err != nil {
		return nil, err
	}

	isEmailInUse := authRepository.IsEmailInUse(payload.Email)

	if isEmailInUse {
		return nil, errors.New(appErrors.EMAIL_IN_USE)
	}

	isNickInUse := usersRepository.IsNickInUse(payload.Nick)

	if isNickInUse {
		return nil, errors.New(appErrors.NICK_IN_USE)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)

	err = authRepository.RegisterUser(payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}