package users

import (
	"core/configs"

	"net/http"
	"strconv"

	"github.com/kuriozapp/toolkit/responses"
)

// SearchUserHandler is a handler to search a user.
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

// GetAuthenticatedUserHandler is a handler retrieve authenticated user data.
//
// It reads user's token, decode it and return the user data.
func GetAuthenticatedUserHandler(handlerCtx *configs.HandlersCtx, usersRepository *UsersRepository) error {
	authenticatedUserId := GetUserByToken(handlerCtx.C).ID

	user, err := GetAuthenticatedUser(handlerCtx, authenticatedUserId, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, user)
}
