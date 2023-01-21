package validations

import (
	helpers "core/src/helpers/requests"
	"core/src/models"
	"core/src/repositories"
	"errors"
)

// IsEmailInUse checks if email is already in use by other user.
func IsEmailInUse(usersRepository *repositories.Users, email string) (bool, models.RequestError) {
	foundUserByEmail := usersRepository.FindUserByEmail(email)

	if foundUserByEmail.Email != "" {
		return true, helpers.NewRequestError(errors.New("o e-mail informado já está em uso"), 403)
	}

	return false, models.RequestError{}
}
