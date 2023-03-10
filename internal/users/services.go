package users

import (
	"core/configs"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	date "github.com/quessapp/toolkit/constants"
	toolkitEntities "github.com/quessapp/toolkit/entities"
)

var (
	// 30 posts/questions by week
	USER_DEFAULT_POST_MONTHLY_LIMIT             = 30
	USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET int64 = 7
)

// SearchUser searches for users based on a search value and returns a paginated list of matching users.
// If the page argument is 0, it sets it to 1 (default). The authenticatedUserId argument is used to filter out the authenticated user from the search results.
// The function returns a pointer to a PaginatedUsers struct representing the paginated list of matching users, and an error, if any occurred during the search process.
func SearchUser(handlerCtx *configs.HandlersCtx, value string, page *int64, authenticatedUserId toolkitEntities.ID, usersRepository *UsersRepository) (*PaginatedUsers, error) {
	if *page == 0 {
		*page = 1
	}

	return usersRepository.Search(value, page)
}

// GetAuthenticatedUser retrieves the authenticated user's data, generates access and refresh tokens for them, and returns a ResponseWithUser struct containing the user's data and tokens.
// The authenticatedUserId argument is used to retrieve the user's data from the usersRepository argument.
// The function returns a pointer to a ResponseWithUser struct representing the user's data and tokens, and an error, if any occurred during the process.
func GetAuthenticatedUser(handlerCtx *configs.HandlersCtx, authenticatedUserId toolkitEntities.ID, usersRepository *UsersRepository) (*ResponseWithUser, error) {
	u := usersRepository.FindUserByID(authenticatedUserId)

	if err := UserExists(u); err != nil {
		return nil, err
	}

	accessToken, err := CreateAccessToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	refreshToken, err := CreateRefreshToken(u, handlerCtx.Cfg)

	if err != nil {
		return nil, err
	}

	user := &ResponseWithUser{
		User: &User{
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

// FindUserByNick searches for a user in the repository by their nickname and returns a pointer to a User struct containing their information.
func FindUserByNick(handlerCtx *configs.HandlersCtx, nick string, usersRepository *UsersRepository) (*User, error) {
	u := usersRepository.FindUserByNick(nick)

	if err := UserExists(u); err != nil {
		return nil, err
	}

	user := &User{
		ID:        u.ID,
		Nick:      u.Nick,
		Name:      u.Name,
		AvatarURL: u.AvatarURL,
	}

	return user, nil
}

// DecrementUserLimit decrements user posts limit.
func DecrementUserLimit(userId toolkitEntities.ID, usersRepository *UsersRepository) error {
	foundUser := usersRepository.FindUserByID(userId)

	if foundUser.IsPRO {
		log.Printf("Not necessary to decrement user %s limit. The user is a PRO member.\n", foundUser.Nick)

		return nil
	}

	foundUser.PostsLimit -= 1

	if err := usersRepository.DecrementLimit(userId, foundUser.PostsLimit); err != nil {
		log.Printf("Fail to decrement user limit %s.\n", err)

		return err
	}

	return nil
}

// UpdateLastPublishedAt updates last publish at fields in database.
func UpdateLastPublishedAt(user *User, usersRepository *UsersRepository) error {
	return usersRepository.UpdateLastPublishedAt(user.ID)
}

// CreateUserToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateUserToken(u *User, expiresIn int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"exp":   expiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// CreateAccessToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateAccessToken(u *User, cfg *configs.Conf) (string, error) {
	return CreateUserToken(u, time.Now().Add(date.ONE_DAY_IN_HOURS).Unix(), cfg.JWTSecret)
}

// CreateRefreshToken creates an user JWT token with followed fields:
// id, name, email, exp. It returns string and error.
func CreateRefreshToken(u *User, cfg *configs.Conf) (string, error) {
	return CreateUserToken(u, time.Now().Add(date.THIRTY_DAYS_IN_HOURS).Unix(), cfg.JWTSecret)
}

// DecodeUserToken decodes an user JWT token with followed fields:
// id, name and email.
func DecodeUserToken(c *fiber.Ctx) toolkitEntities.DecodeUserTokenResult {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	parsedId, _ := toolkitEntities.ParseID(claims["id"].(string))

	u := toolkitEntities.DecodeUserTokenResult{
		ID:    parsedId,
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}

	return u
}

// GetUserByToken decodes a token and get user info from token.
func GetUserByToken(c *fiber.Ctx) toolkitEntities.DecodeUserTokenResult {
	return DecodeUserToken(c)
}

// ResetLimit resets user limit
func ResetLimit(u *User, usersRepository *UsersRepository) error {
	currentDate := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		time.Now().Nanosecond(),
		time.UTC,
	)

	lastPublish := time.Date(
		u.LastPublishAt.Year(),
		u.LastPublishAt.Month(),
		u.LastPublishAt.Day(),
		u.LastPublishAt.Hour(),
		u.LastPublishAt.Minute(),
		u.LastPublishAt.Second(),
		u.LastPublishAt.Nanosecond(),
		time.UTC,
	)

	diffBetweenLastPublishedAndCurrentDate := currentDate.Sub(lastPublish)
	diffInDays := int64(diffBetweenLastPublishedAndCurrentDate.Hours() / 24)
	canReset := diffInDays >= USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET

	if !canReset {
		log.Printf("It's not necessary to reset limit for user %s because it has not passed %d days since their last publish. Their current limit is %d", u.Nick, USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET, u.PostsLimit)
		return nil
	}

	return usersRepository.DecrementLimit(u.ID, USER_DEFAULT_POST_MONTHLY_LIMIT)
}
