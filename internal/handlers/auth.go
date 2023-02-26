package handlers

import (
	"core/internal/configs"
	"core/internal/dtos"
	"core/internal/services"
	"net/http"

	"github.com/kuriozapp/toolkit/responses"
)

// SignUpUserHandler is a handler to sign up an user.
func SignUpUserHandler(handlerCtx *configs.HandlersCtx) error {
	payload := dtos.SignUpUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignUp(handlerCtx, &payload)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}

// SignInUserHandler is a handler to sign in an user.
func SignInUserHandler(handlerCtx *configs.HandlersCtx) error {
	payload := dtos.SignInUserDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	u, err := services.SignIn(handlerCtx, &payload)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, u)
}
