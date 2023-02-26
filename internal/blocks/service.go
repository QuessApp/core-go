package blocks

import (
	"core/internal/users"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// BlockUser blocks an user.
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
