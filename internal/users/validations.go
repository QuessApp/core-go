package users

import (
	"errors"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// UserExists checks if the given user exists by verifying if the ID field is non-zero.
// If the ID is zero, an error with the message "USER_NOT_FOUND" is returned.
// Otherwise, nil is returned, indicating that the user exists.
func UserExists(u *User) error {
	if toolkitEntities.IsZeroID(u.ID) {
		return errors.New(USER_NOT_FOUND)
	}

	return nil
}
