package services

import (
	"core/internal/configs"
	"core/internal/dtos"
	"core/internal/entities"
	validations "core/internal/validations/services"
	"core/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

// SignUp reads payload from request body then try to register a new user in database.
func SignUp(handlerCtx *configs.HandlersCtx, payload *dtos.SignUpUserDTO) (*entities.ResponseWithUser, error) {
	payload.Format()

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	if err := validations.IsEmailInUse(handlerCtx.AuthRepository.IsEmailInUse(payload.Email)); err != nil {
		return nil, err
	}

	if err := validations.IsNickInUse(handlerCtx.UsersRepository.IsNickInUse(payload.Nick)); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	payload.Password = string(hashedPassword)

	u, err := handlerCtx.AuthRepository.SignUp(payload)

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

	data := &entities.ResponseWithUser{
		User: &entities.User{
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

// SignIn reads nick and password from an user and try to return user's data.
func SignIn(handlerCtx *configs.HandlersCtx, payload *dtos.SignInUserDTO) (*entities.ResponseWithUser, error) {
	u := handlerCtx.UsersRepository.FindUserByNick(payload.Nick)

	if err := validations.UserExists(u); err != nil {
		return nil, err
	}

	if err := validations.IsPasswordCorrect(bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password))); err != nil {
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

	data := &entities.ResponseWithUser{
		User: &entities.User{
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
