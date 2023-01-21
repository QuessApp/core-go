package validations

import (
	"core/src/models"
	"core/src/repositories"
)

// IsNickInUse checks if nick is already in use by other user.
func IsNickInUse(usersRepository *repositories.Users, nick string) (bool, models.RequestError) {
	foundUserByNick := usersRepository.FindUserByNick(nick)

	if foundUserByNick.Nick != "" {
		return true, models.RequestError{
			Message: "O nick informado já está em uso.",
			Status:  403,
		}
	}

	return false, models.RequestError{}
}
