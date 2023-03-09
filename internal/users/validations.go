package users

import (
	"errors"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// UserExists returns error message if user does not exists.
func UserExists(u *User) error {
	if toolkitEntities.IsZeroID(u.ID) {
		return errors.New(USER_NOT_FOUND)
	}

	return nil
}
