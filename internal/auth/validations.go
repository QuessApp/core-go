package auth

import (
	"errors"
)

// IsEmailInUse returns error if provided email is already in use.
func IsEmailInUse(isEmailInUse bool) error {
	if isEmailInUse {
		return errors.New(EMAIL_IN_USE)
	}

	return nil
}

// IsNickInUse returns error if provided nick is already in use.
func IsNickInUse(isNickInUse bool) error {
	if isNickInUse {
		return errors.New(NICK_IN_USE)
	}

	return nil
}

// IsPasswordCorrect returns an error if hashed password don't match.
func IsPasswordCorrect(hashResult error) error {
	if hashResult != nil {
		return errors.New(INCORRECT_SIGNIN_DATA)
	}

	return nil
}