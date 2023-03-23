package auth

import (
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/emails"
	"github.com/quessapp/core-go/internal/users"
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"golang.org/x/crypto/bcrypt"
)

// SignUp is a function for signing up a user. It takes in several parameters, including a HandlersCtx struct, a SignUpUserDTO payload, an AuthRepository, and a UsersRepository.
// The function first formats the payload using the Format() method defined in the SignUpUserDTO struct. It then validates the payload using the Validate() method also defined in the SignUpUserDTO struct.
// Next, the function checks if the email and nick are already in use using the IsEmailInUse() and IsNickInUse() methods defined in the users package.
// If the payload is valid and the email and nick are not already in use, the function generates a hashed password using the bcrypt package and the payload's password.
// The function then calls the SignUp() method of the AuthRepository and passes in the payload. If the signup is successful,
// the function creates an access token and refresh token for the user using the CreateAccessToken() and CreateRefreshToken() methods defined in the users package.
// Finally, the function creates a ResponseWithUser struct containing the user's ID, name, email, locale, access token, and refresh token, and returns it along with any error that occurred during the process.
func SignUp(handlerCtx *configs.HandlersCtx, payload *SignUpUserDTO, authRepository *AuthRepository, usersRepository *users.UsersRepository) (*users.ResponseWithUser, error) {
	payload.Format()

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if err := users.IsEmailInUse(usersRepository.IsEmailInUse(payload.Email)); err != nil {
		return nil, err
	}

	if err := users.IsNickInUse(usersRepository.IsNickInUse(payload.Nick)); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)

	u, err := authRepository.SignUp(payload)

	if err != nil {
		return nil, err
	}

	authTokens, err := authRepository.CreateAuthTokens(u.ID, handlerCtx.Cfg.JWTSecret)

	if err != nil {
		return nil, err
	}

	data := &users.ResponseWithUser{
		User: &users.User{
			ID:        u.ID,
			AvatarURL: u.AvatarURL,
			Name:      u.Name,
			Email:     u.Email,
			Locale:    u.Locale,
		},
		AccessToken:  authTokens.AccessToken,
		RefreshToken: authTokens.RefreshToken,
	}

	return data, nil
}

// SignIn function is responsible for authenticating a user with their nickname and password.
//
// It receives a HandlersCtx struct containing the configuration of the app's handlers,
// a pointer to a SignInUserDTO struct containing the user's nickname and password,
// and a pointer to a UsersRepository struct responsible for accessing user data in the database.
//
// It returns a ResponseWithUser struct containing the authenticated user's information,
// an access token and a refresh token if the authentication was successful.
// Otherwise, it returns an error.
func SignIn(handlerCtx *configs.HandlersCtx, payload *SignInUserDTO, authRepository *AuthRepository, usersRepository *users.UsersRepository) (*users.ResponseWithUser, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}

	u := usersRepository.FindUserByNick(payload.Nick)

	if err := users.UserExists(u); err != nil {
		return nil, err
	}

	if err := IsPasswordCorrect(bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password))); err != nil {
		return nil, err
	}

	authTokens, err := authRepository.CreateAuthTokens(u.ID, handlerCtx.Cfg.JWTSecret)

	if err != nil {
		return nil, err
	}

	data := &users.ResponseWithUser{
		User: &users.User{
			ID:              u.ID,
			AvatarURL:       u.AvatarURL,
			Name:            u.Name,
			Email:           u.Email,
			PostsLimit:      u.PostsLimit,
			IsPRO:           u.IsPRO,
			EnableAPPEmails: u.EnableAPPEmails,
			Locale:          u.Locale,
		},
		AccessToken:  authTokens.AccessToken,
		RefreshToken: authTokens.RefreshToken,
	}

	return data, nil
}

// RefreshToken generates a new access token and refresh token pair for the authenticated user
// identified by the given userID and refresh token. It first checks if the token exists in the
// database using the AuthRepository's FindTokenByUserIDAndRefreshToken function. If the token
// doesn't exist, it returns an error. Otherwise, it deletes the existing token using the
// AuthRepository's DeleteByID function and creates a new token pair using the CreateAuthTokens
// function. It returns the new token pair or an error if there was an issue.
func RefreshToken(handlerCtx *configs.HandlersCtx, authenticatedUserID toolkitEntities.ID, refreshToken string, authRepository *AuthRepository) (*Token, error) {
	t := authRepository.FindTokenByUserIDAndRefreshToken(authenticatedUserID, refreshToken)

	if err := IsTokenExpired(t); err != nil {
		return nil, err
	}

	if err := TokenExists(t); err != nil {
		return nil, err
	}

	err := authRepository.DeleteByID(t.ID)

	if err != nil {
		return nil, err
	}

	return authRepository.CreateAuthTokens(*t.CreatedBy, handlerCtx.Cfg.JWTSecret)
}

// Logout deletes all the tokens associated with the authenticated user identified by the given userID.
// It returns an error if there was an issue.
// Otherwise, it returns nil.
func Logout(handlerCtx *configs.HandlersCtx, authenticatedUserID toolkitEntities.ID, authRepository *AuthRepository) error {
	tokenType := "Bearer"
	err := authRepository.DeleteAllUserTokens(authenticatedUserID, &tokenType)

	if err != nil {
		return err
	}

	return nil
}

// ForgotPassword function handles the password reset process.
// It takes a HandlersCtx, a ForgotPasswordDTO, an AuthRepository, and a UsersRepository as arguments.
// The function first finds the user with the given email address using the UsersRepository.
// If the user does not exist, it returns an error.
// If the user exists, the function deletes all of the user's tokens of type "Code" using the AuthRepository.
// Then, it creates a new code token for the user using the AuthRepository.
// Finally, it sends an email to the user with the code token using the EmailsQueue and the Emails package.
// If any error occurs, the function returns the error. Otherwise, it returns nil.
func ForgotPassword(handlerCtx *configs.HandlersCtx, payload ForgotPasswordDTO, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	u := usersRepository.FindUserByEmail(payload.Email)

	if err := users.UserExists(u); err != nil {
		return err
	}

	tokenType := "Code"
	err := authRepository.DeleteAllUserTokens(u.ID, &tokenType)

	if err != nil {
		return err
	}

	t, err := authRepository.CreateCodeToken(u.ID)

	if err != nil {
		return err
	}

	if err := emails.SendEmailForgotPassword(handlerCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.EmailsQueue, t.Code, u); err != nil {
		return err
	}

	return nil
}
