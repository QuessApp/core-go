package auth

import (
	"context"
	"errors"

	"golang.org/x/oauth2"

	"github.com/quessapp/core-go/configs"

	"github.com/quessapp/core-go/internal/users"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(handlerCtx *configs.HandlersCtx, code string) (*ResponseWithAuthenticatedUserData, error) {
	t, err := handlerCtx.OAuth.Exchange(context.Background(), code)

	if err != nil {
		return nil, err
	}

	IDToken, ok := t.Extra("id_token").(string)

	if !ok {
		return nil, errors.New("foo")
	}

	userInfo, err := handlerCtx.OpenIDClient.UserInfo(context.Background(), oauth2.StaticTokenSource(t))

	if err != nil {
		return nil, err
	}

	data := &ResponseWithAuthenticatedUserData{
		Tokens: Tokens{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			Expiry:       t.Expiry.String(),
			TokenType:    t.TokenType,
			IDToken:      IDToken,
		},
		User: userInfo,
	}

	return data, nil
}

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

	accessToken, err := users.CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := users.CreateRefreshToken(u, handlerCtx.Cfg)

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
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
func SignIn(handlerCtx *configs.HandlersCtx, payload *SignInUserDTO, usersRepository *users.UsersRepository) (*users.ResponseWithUser, error) {
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

	accessToken, err := users.CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := users.CreateRefreshToken(u, handlerCtx.Cfg)

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
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return data, nil
}
