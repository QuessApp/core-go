package validations

import (
	pkg "core/pkg/errors"
	"errors"
)

// DidBlockedReceiver returns error message if user is blocked or catch any error.
func DidBlockedReceiver(didBlockedReceiver bool, err error) error {
	if err != nil {
		return err
	}

	if didBlockedReceiver {
		return errors.New(pkg.DID_BLOCKED_RECEIVER)
	}

	return nil
}

// IsBlockedByReceiver returns error message if user is blocked by receiver of the question.
func IsBlockedByReceiver(isBlockedByReceiver bool, err error) error {
	if err != nil {
		return err
	}

	if isBlockedByReceiver {
		return errors.New(pkg.BLOCKED_BY_RECEIVER)
	}

	return nil
}
