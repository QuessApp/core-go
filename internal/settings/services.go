package settings

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

// UpdatePreferences updates a user's preferences based on the data provided in the UpdatePreferencesDTO object.
// This function takes a HandlersCtx object, a pointer to an UpdatePreferencesDTO object, the authenticated user's ID, and a UsersRepository object as parameters.
// It updates the user's preferences using the data in the UpdatePreferencesDTO object by calling the UpdatePreferences method of the UsersRepository object.
// If any error occurs during this process, it returns that error.
func UpdatePreferences(handlerCtx *configs.HandlersCtx, payload *users.UpdatePreferencesDTO, authenticatedUserID toolkitEntities.ID, usersRepository *users.UsersRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	return usersRepository.UpdatePreferences(authenticatedUserID, payload)
}
