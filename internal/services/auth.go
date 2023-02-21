package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	appErrors "core/pkg/errors"

	"errors"

	"golang.org/x/crypto/bcrypt"
)

// SignUp reads payload from request body then try to register a new user in database.
func SignUp(payload *entities.User, usersRepository *repositories.Users, authRepository *repositories.Auth) (*entities.User, error) {
	payload.Format()

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if isEmailInUse := authRepository.IsEmailInUse(payload.Email); isEmailInUse {
		return nil, errors.New(appErrors.EMAIL_IN_USE)
	}

	if isNickInUse := usersRepository.IsNickInUse(payload.Nick); isNickInUse {
		return nil, errors.New(appErrors.NICK_IN_USE)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)

	if err = authRepository.SignUp(payload); err != nil {
		return nil, err
	}

	return payload, nil
}
