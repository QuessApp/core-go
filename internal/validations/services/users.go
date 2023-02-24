package validations

import (
	"core/internal/entities"
	pkg "core/pkg/errors"
	"errors"
)

// UserExists returns error message if user does not exists.
func UserExists(u *entities.User) error {
	if u == nil {
		return errors.New(pkg.USER_NOT_FOUND)
	}

	return nil
}
