package users

import (
	"core/configs"

	"net/http"
	"strconv"

	"github.com/quessapp/toolkit/responses"
)

// SearchUsersByValue performs a search for users based on a search value.
// It returns a list of users matching the search, if any. The page of search results can be specified using the "page" parameter.
// The authenticated user ID is obtained from the JWT token in the request context.
// If an error occurs during the search or parsing of parameters, a Bad Request response is returned.
// Otherwise, a successful response is returned with the list of matching users.
func SearchUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *UsersRepository) error {
	value := handlerCtx.C.Query("search")

	p, err := strconv.Atoi(handlerCtx.C.Query("page"))

	page := int64(p)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	authenticatedUserId := GetUserByToken(handlerCtx.C).ID

	users, err := SearchUser(handlerCtx, value, &page, authenticatedUserId, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, users)
}

// GetAuthenticatedUserHandler retrieves the authenticated user from the database based on the authenticated user ID and returns a HTTP response.
// It returns a successful HTTP response with the authenticated user's data if the user is found.
// Otherwise, it returns a Bad Request HTTP response with an error message.
func GetAuthenticatedUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *UsersRepository) error {
	authenticatedUserId := GetUserByToken(handlerCtx.C).ID

	user, err := GetAuthenticatedUser(handlerCtx, authenticatedUserId, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, user)
}

// FindUserByNickHandler retrieves a user from the database based on their nickname and returns a HTTP response.
// It takes the nickname as a parameter from the request context and returns a successful HTTP response with the user's data if the user is found.
// Otherwise, it returns a Bad Request HTTP response with an error message.
func FindUserByNickHandler(handlerCtx *configs.HandlersCtx, usersRepository *UsersRepository) error {
	nick := handlerCtx.C.Params("nick")

	user, err := FindUserByNick(handlerCtx, nick, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, user)
}
