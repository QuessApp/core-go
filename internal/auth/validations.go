package auth

import (
	"errors"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"
)

// IsPasswordCorrect returns an error if hashed password don't match.
func IsPasswordCorrect(hashResult error) error {
	if hashResult != nil {
		return errors.New(pkgErrors.INCORRECT_SIGNIN_DATA)
	}

	return nil
}
