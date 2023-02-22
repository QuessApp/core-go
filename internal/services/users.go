package services

import (
	"core/internal/repositories"
	pkg "core/pkg/entities"
	"log"
)

// DecrementUserLimit decrements user posts limit.
func DecrementUserLimit(userId pkg.ID, usersRepository *repositories.Users) error {
	foundUser, err := usersRepository.FindUserByID(userId)

	log.Printf("Fail to find user by id %s\n", err)

	if err != nil {
		return err
	}

	if foundUser.IsPRO {
		log.Printf("Not necessary to decrement user %s limit. The user is a PRO member.\n", foundUser.Nick)

		return nil
	}

	foundUser.PostsLimit -= 1

	if err := usersRepository.DecrementLimit(userId, foundUser.PostsLimit); err != nil {
		log.Printf("Fail to decrement user limit %s.\n", err)

		return err
	}

	return nil
}
