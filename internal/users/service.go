package users

import (
	"log"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// DecrementUserLimit decrements user posts limit.
func DecrementUserLimit(userId toolkitEntities.ID, usersRepository *UsersRepository) error {
	foundUser := usersRepository.FindUserByID(userId)

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
