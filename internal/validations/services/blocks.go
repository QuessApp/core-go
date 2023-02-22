package validations

import (
	pkg "core/pkg/errors"
	"errors"
)

// ValidateDidBlockedReceiver returns error message if user is blocked or catch any error.
func ValidateDidBlockedReceiver(didBlockedReceiver bool, err error) error {
	if err != nil {
		return err
	}

	if didBlockedReceiver {
		return errors.New(pkg.DID_BLOCKED_RECEIVER)
	}

	return nil
}

// ValidateIsBlockedByReceiver returns error message if user is blocked by receiver of the question.
func ValidateIsBlockedByReceiver(isBlockedByReceiver bool, err error) error {
	if err != nil {
		return err
	}

	if isBlockedByReceiver {
		return errors.New(pkg.BLOCKED_BY_RECEIVER)
	}

	return nil
}
