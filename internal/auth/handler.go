package auth

import (
	"core/configs"
	"core/internal/users"
	"net/http"

	"github.com/kuriozapp/toolkit/responses"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	payload := SignUpUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := SignUp(handlerCtx, &payload, authRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}

// SignInUserHandler is a handler to sign in an user.
func SignInUserHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	payload := SignInUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := SignIn(handlerCtx, &payload, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}
