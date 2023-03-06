package users

import (
	"errors"
)

// UserExists returns error message if user does not exists.
func UserExists(u *User) error {
	if u == nil {
		return errors.New(USER_NOT_FOUND)
	}

	return nil
}
