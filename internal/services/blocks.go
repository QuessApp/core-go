package services

import (
	"core/internal/entities"
	"core/internal/repositories"
	validations "core/internal/validations/services"
)

// BlockUser blocks an user.
func BlockUser(payload *entities.BlockedUser, usersRepository *repositories.Users, blocksRepository *repositories.Blocks) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	doesUserToBeBlockedExists, err := usersRepository.FindUserByID(payload.UserToBlock)

	if err != nil {
		return err
	}

	if err := validations.UserExists(doesUserToBeBlockedExists); err != nil {
		return err
	}

	if err := blocksRepository.BlockUser(payload); err != nil {

		return err
	}

	return nil
}
