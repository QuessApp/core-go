package blocks

import (
	"github.com/quessapp/core-go/internal/users"

	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// BlockUser is a function that blocks a user given a payload, a UsersRepository, and a BlocksRepository.
// It takes in a payload of type *BlockUserDTO, a UsersRepository, and a BlocksRepository, and returns an error.
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

// UnblockUser is a function that unblocks a user given a userID, a UsersRepository, and a BlocksRepository.
// It takes in a userID of type toolkitEntities.ID, a UsersRepository, and a BlocksRepository, and returns an error.
func UnblockUser(userID toolkitEntities.ID, usersRepository *users.UsersRepository, blocksRepository *BlocksRepository) error {
	if err := IsReallyBlocked(blocksRepository.IsUserBlocked(userID)); err != nil {
		return err
	}

	err := blocksRepository.UnblockUser(userID)

	if err != nil {
		return err
	}

	return nil
}
