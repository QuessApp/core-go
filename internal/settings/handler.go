package settings

import (
	"core/configs"
	"core/internal/users"
	"net/http"

	"github.com/kuriozapp/toolkit/responses"
)

// UpdatePreferencesHandler is a handler to update user preferences.
func UpdatePreferencesHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository) error {
	payload := users.UpdatePreferencesDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := users.GetUserByToken(handlerCtx.C).ID

	if err := UpdatePreferences(handlerCtx, &payload, authenticatedUserId, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
