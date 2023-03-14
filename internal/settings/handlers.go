package settings

import (
	"net/http"

	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/users"

	"github.com/quessapp/toolkit/responses"
)

// UpdatePreferencesHandler is an HTTP request handler function that updates a user's preferences based on the data provided in the request body.
// This function takes a HandlersCtx object and a UsersRepository object as parameters.
// It attempts to parse the request body into an UpdatePreferencesDTO object, and returns an error response with a 400 Bad Request status code
// if the parsing fails. It then gets the authenticated user's ID from the request context, and calls the UpdatePreferences function to update
// the user's preferences. If any error occurs during this process, it returns an error response with a 400 Bad Request status code.
// If the update is successful, it returns a successful response with a 200 OK status code.
func UpdatePreferencesHandler(handlerCtx *configs.HandlersCtx, usersRepository *users.UsersRepository) error {
	payload := users.UpdatePreferencesDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserID := users.GetUserByToken(handlerCtx.C).ID

	if err := UpdatePreferences(handlerCtx, &payload, authenticatedUserID, usersRepository); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
