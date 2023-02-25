package validations

import (
	"core/internal/entities"
	internalErrors "core/internal/errors"
	"errors"
)

// UserExists returns error message if user does not exists.
func UserExists(u *entities.User) error {
	if u == nil {
		return errors.New(internalErrors.USER_NOT_FOUND)
	}

	return nil
}
