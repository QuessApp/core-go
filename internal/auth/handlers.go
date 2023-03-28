package auth

import (
	"strings"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/users"
	"github.com/quessapp/toolkit/responses"

	"net/http"
)

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

	u, err := SignIn(handlerCtx, &payload, authRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, u)
}

// RefreshTokenHandler handles the incoming HTTP request for refreshing a user's token.
// It extracts the refresh token from the incoming request's Authorization header and uses it
// to retrieve the authenticated user's ID. It then calls the RefreshToken function to generate
// a new access token and refresh token pair. If RefreshToken returns an error, it returns a
// BadRequest response. Otherwise, it returns a Success response containing the new token pair.
func RefreshTokenHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	h := handlerCtx.C.Get("Authorization")

	refreshToken := strings.Split(string(h), "Bearer ")[1]

	authenticatedUserID := users.GetUserByToken(handlerCtx).ID

	t, err := RefreshToken(handlerCtx, authenticatedUserID, refreshToken, authRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, t)
}

// LogoutHandler handles the incoming HTTP request for logging out a user.
// It extracts the refresh token from the incoming request's Authorization header and uses it
// to retrieve the authenticated user's ID. It then calls the Logout function to invalidate
// the user's refresh token. If Logout returns an error, it returns a BadRequest response.
// Otherwise, it returns a Success response.
func LogoutHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	authenticatedUserID := users.GetUserByToken(handlerCtx).ID
	oldToken := strings.Split(handlerCtx.C.Get("Authorization"), "Bearer ")[1]

	err := Logout(handlerCtx, authenticatedUserID, oldToken, authRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}

// ForgotPasswordHandler function is an HTTP request handler that handles the password reset process.
// It takes a HandlersCtx, an AuthRepository, and a UsersRepository as arguments.
// First, it extracts the ForgotPasswordDTO from the request body using the HandlersCtx.
// Then, it calls the ForgotPassword function passing the extracted DTO and repositories.
// If the ForgotPassword function returns an error, the function returns an HTTP response with a status code of 400 and the error message.
// If the ForgotPassword function does not return an error, the function returns an HTTP response with a status code of 200 and a null body.
func ForgotPasswordHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	payload := ForgotPasswordDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	err := ForgotPassword(handlerCtx, payload, authRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusOK, nil)
}

// ResetPasswordHandler function is an HTTP request handler that handles the password reset process.
// It takes a HandlersCtx, an AuthRepository, and a UsersRepository as arguments.
// First, it extracts the ResetPasswordDTO from the request body using the HandlersCtx.
// Then, it calls the ResetPassword function passing the extracted DTO and repositories.
// If the ResetPassword function returns an error, the function returns an HTTP response with a status code of 400 and the error message.
// If the ResetPassword function does not return an error, the function returns an HTTP response with a status code of 201 and a null body.
func ResetPasswordHandler(handlerCtx *configs.HandlersCtx, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	payload := ResetPasswordDTO{}

	if err := handlerCtx.C.BodyParser(&payload); err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	err := ResetPassword(handlerCtx, payload, authRepository, usersRepository)

	if err != nil {
		return responses.ParseUnsuccesfull(handlerCtx.C, http.StatusBadRequest, err.Error())
	}

	return responses.ParseSuccessful(handlerCtx.C, http.StatusCreated, nil)
}
