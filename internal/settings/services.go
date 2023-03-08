package settings

import (
	"core/configs"
	"core/internal/users"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// UpdatePreferences updates user preferences such as emails, etc.
func UpdatePreferences(handlerCtx *configs.HandlersCtx, payload *users.UpdatePreferencesDTO, authenticatedUserId toolkitEntities.ID, usersRepository *users.UsersRepository) error {
	return usersRepository.UpdatePreferences(authenticatedUserId, payload)
}
