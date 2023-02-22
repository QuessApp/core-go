package validations

import (
	pkg "core/pkg/errors"
	"errors"
)

// ValidateIsEmailInUse returns error if provided email is already in use.
func ValidateIsEmailInUse(isEmailInUse bool) error {
	if isEmailInUse {
		return errors.New(pkg.EMAIL_IN_USE)
	}

	return nil
}

// ValidateIsNickInUse returns error if provided nick is already in use.
func ValidateIsNickInUse(isNickInUse bool) error {
	if isNickInUse {
		return errors.New(pkg.NICK_IN_USE)
	}

	return nil
}
