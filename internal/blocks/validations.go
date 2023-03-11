package blocks

import (
	pkgErrors "core/pkg/errors"
	"errors"
)

// DidBlockedReceiver returns error message if user is blocked or catch any error.
func DidBlockedReceiver(didBlockedReceiver bool) error {
	if didBlockedReceiver {
		return errors.New(pkgErrors.DID_BLOCKED_RECEIVER)
	}

	return nil
}

// IsBlockedByReceiver returns error message if user is blocked by receiver of the question.
func IsBlockedByReceiver(isBlockedByReceiver bool) error {
	if isBlockedByReceiver {
		return errors.New(pkgErrors.BLOCKED_BY_RECEIVER)
	}

	return nil
}

// IsAlreadyBlocked returns error message if user is already blocked.
func IsAlreadyBlocked(isUserAlreadyBlocked bool) error {
	if isUserAlreadyBlocked {
		return errors.New(pkgErrors.ALREADY_BLOCKED)
	}

	return nil
}

// IsBlockingYourself returns error message if user is trying to block yourself.
func IsBlockingYourself(payload *BlockUserDTO) error {
	if payload.BlockedBy == payload.UserToBlock {
		return errors.New(pkgErrors.CANT_BLOCK_YOURSELF)
	}

	return nil
}

// IsReallyBlocked returns error message if user try to unblock an user but the user is not really blocked.
func IsReallyBlocked(isUserBlocked bool) error {
	if !isUserBlocked {
		return errors.New(pkgErrors.CANT_UNBLOCK_NOT_BLOCKED)
	}

	return nil
}
