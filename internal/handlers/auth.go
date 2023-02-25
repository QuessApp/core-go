package handlers

import (
	"core/cmd/app/entities"
	"core/internal/dtos"
	"core/internal/services"
	"core/pkg/responses"
	"net/http"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(handlerCtx *entities.HandlersContext) error {
	payload := dtos.SignUpUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignUp(handlerCtx.Cfg, &payload, handlerCtx.UsersRepository, handlerCtx.AuthRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}

// SignInUserHandler is a handler to sign in an user.
func SignInUserHandler(handlerCtx *entities.HandlersContext) error {
	payload := dtos.SignInUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignIn(handlerCtx.Cfg, &payload, handlerCtx.UsersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}
