package validations

import (
	"core/src/models"
	"core/src/repositories"
)

// IsEmailInUse checks if email is already in use by other user.
func IsEmailInUse(usersRepository *repositories.Users, email string) (bool, models.RequestError) {
	foundUserByEmail := usersRepository.FindUserByEmail(email)

	if foundUserByEmail.Email != "" {
		return true, models.RequestError{
			Message: "O e-mail informado já está em uso.",
			Status:  403,
		}
	}

	return false, models.RequestError{}
}
