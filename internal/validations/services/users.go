package validations

import (
	"core/internal/entities"
	pkg "core/pkg/errors"
	"errors"
)

// UserExists returns error message if user does not exists.
func UserExists(u *entities.User) error {
	userExists := u.Nick != ""

	if !userExists {
		return errors.New(pkg.USER_NOT_FOUND)
	}

	return nil
}
