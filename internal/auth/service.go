package auth

import (
	"core/configs"

	"core/internal/users"
	"core/pkg/jwt"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"
	"golang.org/x/crypto/bcrypt"
)

// SignUp reads payload from request body then try to register a new user in database.
//
// It formats the payload, check if provided email and nick are already in use.
//
// After validations, the provided password will be hashed. If no error is returned, the user will be created in database.
func SignUp(handlerCtx *configs.HandlersCtx, payload *SignUpUserDTO, authRepository *AuthRepository, usersRepository *users.UsersRepository) (*users.ResponseWithUser, error) {
	payload.Format()

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if err := IsEmailInUse(authRepository.IsEmailInUse(payload.Email)); err != nil {
		return nil, err
	}

	if err := IsNickInUse(usersRepository.IsNickInUse(payload.Nick)); err != nil {
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

	accessToken, err := jwt.CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateRefreshToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	data := &users.ResponseWithUser{
		User: &users.User{
			ID:        u.ID,
			AvatarURL: u.AvatarURL,
			Name:      u.Name,
			Email:     u.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return data, nil
}

// SignIn reads nick and password from the request and will try to return user's data.
//
// It will verify if user already exists and after validations the provided password will be hashed and compared with password in database.
//
// If no error is returned, the user will be created in database and access & refresh token will be returned.
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

	accessToken, err := jwt.CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateRefreshToken(u, handlerCtx.Cfg)

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
			EnableAppEmails: u.EnableAppEmails,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return data, nil
}

// GetAuthenticatedUser gets an user from their token.
func GetAuthenticatedUser(handlerCtx *configs.HandlersCtx, userId toolkitEntities.ID, usersRepository *users.UsersRepository) (*users.ResponseWithUser, error) {
	u := usersRepository.FindUserByID(userId)

	if err := users.UserExists(u); err != nil {
		return nil, err
	}

	accessToken, err := jwt.CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.CreateRefreshToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	user := &users.ResponseWithUser{
		User: &users.User{
			ID:         u.ID,
			Nick:       u.Nick,
			Name:       u.Name,
			AvatarURL:  u.AvatarURL,
			Email:      u.Email,
			IsPRO:      u.IsPRO,
			PostsLimit: u.PostsLimit,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return user, nil
}
