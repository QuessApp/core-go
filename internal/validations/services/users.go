package validations

import (
	"core/internal/entities"
	pkg "core/pkg/errors"
	"errors"
)

// ValidateUserExists returns error message if user does not exists.
func ValidateUserExists(user *entities.User) error {
	userExists := user.Nick != ""

	if !userExists {
		return errors.New(pkg.USER_NOT_FOUND)
	}

	return nil
}
