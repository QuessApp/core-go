package auth

import (
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/internal/queues/emails"
	trustedIPs "github.com/quessapp/core-go/internal/queues/trusted-ips"
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

	ip := handlerCtx.C.IP()

	if err := authRepository.AddNewTrustedIPIfDontExists(u.ID, ip); err != nil {
		log.Printf("Error adding new trusted IP: %v for user %v-%v", err, u.ID, u.Nick)
	}

	authTokens, err := authRepository.CreateAuthTokens(u.ID, handlerCtx.Cfg.JWT.Secret)

	if err != nil {
		return nil, err
	}

	data := &users.ResponseWithUser{
		User: &users.User{
			ID:        u.ID,
			Nick:      u.Nick,
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

	ip := handlerCtx.C.IP()
	u := usersRepository.FindUserByNick(payload.Nick)

	if !authRepository.CheckIfTrustedIPExists(u.ID, ip) {
		log.Printf("IP %s is not trusted \n", ip)
		trustedIPs.SendIPToQueue(handlerCtx.Cfg, handlerCtx.MessageQueueCh, handlerCtx.TrustedIPsQueue, ip, u.Email)
	}

	if err := users.UserExists(u); err != nil {
		return nil, err
	}

	if err := IsPasswordCorrect(bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password))); err != nil {
		return nil, err
	}

	authTokens, err := authRepository.CreateAuthTokens(u.ID, handlerCtx.Cfg.JWT.Secret)

	if err != nil {
		return nil, err
	}

	if payload.TrustIP {
		if err := authRepository.AddNewTrustedIPIfDontExists(u.ID, ip); err != nil {
			log.Printf("Error adding new trusted IP: %v for user %v-%v", err, u.ID, u.Nick)
		}
	}

	data := &users.ResponseWithUser{
		User: &users.User{
			ID:              u.ID,
			AvatarURL:       u.AvatarURL,
			Nick:            u.Nick,
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
		if err := authRepository.DeleteTokenByID(t.ID); err != nil {
			return nil, err
		}

		return nil, err
	}

	if err := TokenExists(t); err != nil {
		return nil, err
	}

	if err := authRepository.DeleteTokenByID(t.ID); err != nil {
		return nil, err
	}

	return authRepository.CreateAuthTokens(*t.CreatedBy, handlerCtx.Cfg.JWT.Secret)
}

// Logout deletes the refresh token from the database.
// It takes a HandlersCtx, an authenticatedUserID, a token, and an AuthRepository as arguments.
// The function first deletes the token from the database using the AuthRepository's DeleteRefreshToken function.
// If any error occurs, the function returns the error. Otherwise, it returns nil.
func Logout(handlerCtx *configs.HandlersCtx, authenticatedUserID toolkitEntities.ID, token string, authRepository *AuthRepository) error {
	return authRepository.DeleteRefreshToken(token)
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
	if err := authRepository.DeleteAllUserTokens(u.ID, &tokenType); err != nil {
		return err
	}

	t, err := authRepository.CreateCodeToken(u.ID)

	if err != nil {
		return err
	}

	if err := emails.SendEmailForgotPassword(handlerCtx, t.Code, u); err != nil {
		return err
	}

	return nil
}

// ResetPassword resets a user's password based on the provided ResetPasswordDTO.
// It takes in the HandlerCtx, ResetPasswordDTO, AuthRepository, and UsersRepository as parameters,
// and returns an error if any of the operations fail.
// The function first checks if the code exists in the database using the AuthRepository's FindTokenByCode function.
// If the code does not exist, it returns an error.
// If the code exists, the function checks if the code is expired using the IsCodeExpired function.
// If the code is expired, it deletes the code token from the database using the AuthRepository's DeleteTokenByID function.
// Then, it returns an error.
// If the code is not expired, the function generates a new hashed password using the bcrypt package.
// Then, it finds the user associated with the code token using the UsersRepository's FindUserByID function.
// If the user does not exist, it returns an error.
// If the user exists, the function updates the user's password using the UsersRepository's UpdateUserPassword function.
// Then, it deletes the code token from the database using the AuthRepository's DeleteTokenByID function.
// Finally, it returns nil.
func ResetPassword(handlerCtx *configs.HandlersCtx, payload ResetPasswordDTO, authRepository *AuthRepository, usersRepository *users.UsersRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	t := authRepository.FindTokenByCode(payload.Code)

	if err := CodeExists(t); err != nil {
		return err
	}

	if err := IsCodeExpired(t); err != nil {
		if err := authRepository.DeleteTokenByID(t.ID); err != nil {
			return err
		}

		return err
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u := usersRepository.FindUserByID(*t.CreatedBy)

	if err := users.UserExists(u); err != nil {
		return err
	}

	if err := authRepository.UpdateUserPassword(u.ID, newHashedPassword); err != nil {
		return err
	}

	if payload.LogoutFromAllDevices {
		// Bearer is the token type used for auth
		// We want to delete all access tokens (logout user)
		// if we remove all tokens, user can not refresh the token
		// and will be logged out from all devices
		tokenType := "Bearer"

		if err := authRepository.DeleteAllUserTokens(u.ID, &tokenType); err != nil {
			return err
		}
	}

	if err := authRepository.DeleteTokenByID(t.ID); err != nil {
		return err
	}

	go emails.SendEmailPasswordChanged(handlerCtx, u)

	return nil
}
