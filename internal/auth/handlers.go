package auth

import (
	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/users"

	"net/http"

	"github.com/quessapp/toolkit/responses"
)

func RedirectToAuthPageHandler(handlerCtx *configs.HandlersCtx) error {
	if err := handlerCtx.C.Redirect(handlerCtx.AppCtx.OAuth.AuthCodeURL("123"), http.StatusTemporaryRedirect); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return nil
}

func AuthenticateHandler(handlerCtx *configs.HandlersCtx) error {
	code := handlerCtx.C.Query("code")

	d, err := AuthenticateUser(handlerCtx, code)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, d)
}

// SignUpUserHandler is an HTTP handler function that handles requests for user sign-up.
// It receives a HandlersCtx containing the HTTP request context, an AuthRepository for authentication,
// and a UsersRepository for user data access. It parses the request body into a SignUpUserDTO,
// creates a new user using the provided AuthRepository and UsersRepository, and returns a JSON response with the created user data.
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

// SignInUserHandler is an HTTP handler function that handles requests for user sign-in.
// It receives a HandlersCtx containing the HTTP request context, an AuthRepository for authentication,
// and a UsersRepository for user data access. It parses the request body into a SignInUserDTO,
// authenticates the user using the provided AuthRepository, and returns a JSON response with the authenticated user data.
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
