package auth

import (
	"errors"
	"time"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// IsPasswordCorrect checks if the given hash result is nil. If it is, it means that the
// password is correct, and it returns nil. Otherwise, it returns an error.
func IsPasswordCorrect(hashResult error) error {
	if hashResult != nil {
		return errors.New(pkgErrors.INCORRECT_SIGNIN_DATA)
	}

	return nil
}

// TokenExists checks if the given token exists in the database. It returns an error if the
// token's ID is zero, indicating that the token does not exist, or nil if the token exists.
func TokenExists(t *Token) error {
	if toolkitEntities.IsZeroID(t.ID) {
		return errors.New(pkgErrors.TOKEN_NOT_FOUND)
	}

	return nil
}

// IsTokenExpired checks if the given token has expired. It returns an error if the token's
// expiration date is before the current date, indicating that the token has expired, or nil
// if the token has not expired.
func IsTokenExpired(t *Token) error {
	if t.ExpiresAt.Before(time.Now()) {
		return errors.New(pkgErrors.TOKEN_EXPIRED)
	}

	return nil
}

// IsCodeExpired checks if the given code has expired. It returns an error if the code's
// expiration date is before the current date, indicating that the code has expired, or nil
// if the code has not expired.
func IsCodeExpired(c *Token) error {
	if c.ExpiresAt.Before(time.Now()) {
		return errors.New(pkgErrors.CODE_EXPIRED)
	}

	return nil
}
