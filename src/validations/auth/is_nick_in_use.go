package validations

import (
	helpers "core/src/helpers/requests"
	"core/src/models"
	"core/src/repositories"
	"errors"
)

// IsNickInUse checks if nick is already in use by other user.
func IsNickInUse(usersRepository *repositories.Users, nick string) (bool, models.RequestError) {
	foundUserByNick := usersRepository.FindUserByNick(nick)

	if foundUserByNick.Nick != "" {
		return true, helpers.NewRequestError(errors.New("o nick informado já está em uso"), 403)
	}

	return false, models.RequestError{}
}
