package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	validations "core/internal/validations/services"

	"golang.org/x/crypto/bcrypt"
)

// SignUp reads payload from request body then try to register a new user in database.
func SignUp(payload *entities.User, usersRepository *repositories.Users, authRepository *repositories.Auth) (*entities.User, error) {
	payload.Format()

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if err := validations.ValidateIsEmailInUse(authRepository.IsEmailInUse(payload.Email)); err != nil {
		return nil, err
	}

	if err := validations.ValidateIsNickInUse(usersRepository.IsNickInUse(payload.Nick)); err != nil {
		return nil, err
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
