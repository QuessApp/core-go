package blocks

import (
	"core/internal/users"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// BlockUser blocks an user.
//
// It formats the payload, checks if user to block exists, already blocked, checks if trying to block yourself.
//
// After validations, if there are not errors, the user will be blocked by other user.
func BlockUser(payload *BlockUserDTO, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	doesUserToBeBlockedExists := usersRepository.FindUserByID(payload.UserToBlock)

	if err := users.UserExists(doesUserToBeBlockedExists); err != nil {
		return err
	}

	if err := IsAlreadyBlocked(blocksRepository.IsUserBlocked(payload.UserToBlock)); err != nil {
		return err
	}

	if err := IsBlockingYourself(payload); err != nil {
		return err
	}

	if err := blocksRepository.BlockUser(payload); err != nil {
		return err
	}

	return nil
}

// UnblockUser unblocks an user.
//
// It checks if user is really blocked.
//
// After validations, if there are not errors, the user will be unblocked by other user.
func UnblockUser(userId toolkitEntities.ID, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	if err := IsReallyBlocked(blocksRepository.IsUserBlocked(userId)); err != nil {
		return err
	}

	err := blocksRepository.UnblockUser(userId)

	if err != nil {
		return err
	}

	return nil
}
