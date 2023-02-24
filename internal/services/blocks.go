package services

import (
	"core/internal/dtos"
	"core/internal/repositories"
	validations "core/internal/validations/services"
	pkg "core/pkg/entities"
)

// BlockUser blocks an user.
func BlockUser(payload *dtos.BlockUserDTO, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	doesUserToBeBlockedExists := usersRepository.FindUserByID(payload.UserToBlock)

	if err := validations.UserExists(doesUserToBeBlockedExists); err != nil {
		return err
	}

	if err := validations.IsAlreadyBlocked(blocksRepository.IsUserBlocked(payload.UserToBlock)); err != nil {
		return err
	}

	if err := validations.IsBlockingYourself(payload); err != nil {
		return err
	}

	if err := blocksRepository.BlockUser(payload); err != nil {
		return err
	}

	return nil
}

// UnblockUser unblocks an user.
func UnblockUser(userId pkg.ID, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) error {
	if err := validations.IsReallyBlocked(blocksRepository.IsUserBlocked(userId)); err != nil {
		return err
	}

	err := blocksRepository.UnblockUser(userId)

	if err != nil {
		return err
	}

	return nil
}
