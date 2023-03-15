package users

import (
	"errors"

	pkgErrors "github.com/quessapp/core-go/pkg/errors"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// UserExists checks if the given user exists by verifying if the ID field is non-zero.
// If the ID is zero, an error with the message "USER_NOT_FOUND" is returned.
// Otherwise, nil is returned, indicating that the user exists.
func UserExists(u *User) error {
	if toolkitEntities.IsZeroID(u.ID) {
		return errors.New(pkgErrors.USER_NOT_FOUND)
	}

	return nil
}

// ReachedMaxSizeLimit checks if the file size is greater than or equal to 1MB.
// It returns an error if the file size exceeds the limit, otherwise it returns nil.
func ReachedMaxSizeLimit(fileSize int64) error {
	isFileSizeGreaterThanOneMB := fileSize > (1024 * 1024)

	if isFileSizeGreaterThanOneMB {
		return errors.New(pkgErrors.MAX_FILE_SIZE)
	}

	return nil
}

// IsAllowedFileType checks if a file type is allowed or not.
// It takes a boolean parameter `isAllowed`, which should be true if the file type is allowed and false otherwise.
// It returns an error if the file type is not allowed, otherwise it returns nil.
func IsAllowedFileType(isAllowed bool) error {
	if !isAllowed {
		return errors.New(pkgErrors.FILE_TYPE_INVALID)
	}

	return nil
}

// IsEmailInUse returns error if provided email is already in use.
func IsEmailInUse(isEmailInUse bool) error {
	if isEmailInUse {
		return errors.New(pkgErrors.EMAIL_IN_USE)
	}

	return nil
}

// IsNickInUse returns error if provided nick is already in use.
func IsNickInUse(isNickInUse bool) error {
	if isNickInUse {
		return errors.New(pkgErrors.NICK_IN_USE)
	}

	return nil
}
