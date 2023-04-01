package users

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/quessapp/core-go/configs"
	toolkitEntities "github.com/quessapp/toolkit/entities"
	toolkitS3 "github.com/quessapp/toolkit/s3"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// 30 posts/questions per week
	USER_DEFAULT_POST_MONTHLY_LIMIT             = 30
	USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET int64 = 7
)

// SearchUser searches for users based on a search value and returns a paginated list of matching users.
// If the page argument is 0, it sets it to 1 (default). The authenticatedUserID argument is used to filter out the authenticated user from the search results.
// The function returns a pointer to a PaginatedUsers struct representing the paginated list of matching users, and an error, if any occurred during the search process.
func SearchUser(handlerCtx *configs.HandlersCtx, value string, page *int64, authenticatedUserID toolkitEntities.ID, usersRepository *UsersRepository) (*PaginatedUsers, error) {
	if *page == 0 {
		*page = 1
	}

	return usersRepository.Search(value, page)
}

// GetAuthenticatedUser retrieves the authenticated user's data and returns a ResponseWithUser struct containing the user's data and tokens.
// The authenticatedUserID argument is used to retrieve the user's data from the usersRepository argument.
// The function returns a pointer to a ResponseWithUser struct representing the user's data and an error, if any occurred during the process.
func GetAuthenticatedUser(handlerCtx *configs.HandlersCtx, authenticatedUserID toolkitEntities.ID, usersRepository *UsersRepository) (*User, error) {
	u := usersRepository.FindUserByID(authenticatedUserID)

	if err := UserExists(u); err != nil {
		return nil, err
	}

	user := &User{
		ID:         u.ID,
		Nick:       u.Nick,
		Name:       u.Name,
		AvatarURL:  u.AvatarURL,
		Email:      u.Email,
		IsPRO:      u.IsPRO,
		PostsLimit: u.PostsLimit,
		Locale:     u.Locale,
	}

	return user, nil
}

// FindUserByNick searches for a user with the given nickname in the given users repository,
// and returns the corresponding User object, if it exists.
// If the user with the given nickname is not found, an error is returned.
// If an error occurs while checking if the user exists, that error is returned as well.
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

// DecrementUserLimit decrements the posts limit of the user with the given ID by one.
// If the user is a PRO member, their limit will not be decremented and no error will be returned.
// If an error occurs while decrementing the limit, that error will be returned.
func DecrementUserLimit(userID toolkitEntities.ID, usersRepository *UsersRepository) error {
	foundUser := usersRepository.FindUserByID(userID)

	if foundUser.IsPRO {
		log.Printf("Not necessary to decrement user %s limit. The user is a PRO member.\n", foundUser.Nick)

		return nil
	}

	foundUser.PostsLimit -= 1

	if err := usersRepository.DecrementLimit(userID, foundUser.PostsLimit); err != nil {
		log.Printf("Fail to decrement user limit %s.\n", err)

		return err
	}

	return nil
}

// DeleteUserAvatar deletes a user's avatar image from S3.
func DeleteUserAvatar(handlerCtx *configs.HandlersCtx, fileName string) error {
	_, err := toolkitS3.DeleteFile(handlerCtx.S3Client, handlerCtx.Cfg.S3.BucketName, fileName)

	return err
}

// UpdateUserAvatar uploads a user's avatar image to S3 and updates the user's
// avatar URL in the database. If the user already has an avatar, the function
// deletes the old avatar from S3 before uploading the new one. The uploaded
// file is given public-read access.
func UpdateUserAvatar(handlerCtx *configs.HandlersCtx, form *multipart.FileHeader, authenticatedUserID toolkitEntities.ID, usersRepository *UsersRepository) error {
	ACL := "public-read"

	u := usersRepository.FindUserByID(authenticatedUserID)

	if err := UserExists(u); err != nil {
		return err
	}

	allowedFileTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if err := IsAllowedFileType(allowedFileTypes[form.Header.Get("Content-Type")]); err != nil {
		return err
	}

	if err := ReachedMaxSizeLimit(form.Size); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s-%s", authenticatedUserID.Hex(), form.Filename)
	fileDir := fmt.Sprintf("./tmp/%s", fileName)

	if err := handlerCtx.C.SaveFile(form, fileDir); err != nil {
		return err
	}

	f, err := os.Open(fileDir)

	if err != nil {
		return err
	}

	defer os.Remove(fileDir)
	defer f.Close()

	newAvatarURI := fmt.Sprintf("%s%s", handlerCtx.Cfg.CDN.URI, fileName)

	if err := usersRepository.UpdateAvatar(authenticatedUserID, newAvatarURI); err != nil {
		return err
	}

	if u.AvatarURL != "" {
		oldAvatarFileName := strings.Split(u.AvatarURL, handlerCtx.Cfg.CDN.URI)

		if len(oldAvatarFileName) < 1 {
			return nil
		}

		log.Printf("deleting user %s old avatar (%s) to upload a new image \n", u.Nick, oldAvatarFileName)

		if err := DeleteUserAvatar(handlerCtx, oldAvatarFileName[1]); err != nil {
			return err
		}
	}

	_, err = toolkitS3.UploadFile(handlerCtx.S3Client, handlerCtx.Cfg.S3.BucketName, fileName, f, &ACL)

	if err != nil {
		return err
	}

	return nil
}

// UpdateLastPublishedAt updates the last published at timestamp for the given user.
// This function takes a pointer to a User object and a UsersRepository object as parameters.
// It updates the last published at timestamp for the user in the repository, and returns any error that may occur.
func UpdateLastPublishedAt(user *User, usersRepository *UsersRepository) error {
	return usersRepository.UpdateLastPublishedAt(user.ID)
}

// ResetLimit checks if the user's posts limit can be reset based on the USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET constant.
// If the user's last publish date is greater than or equal to USER_POST_MONTHLY_LIMIT_DAYS_TO_RESET days in the past,
// their posts limit will be reset to the default value specified in the USER_DEFAULT_POST_MONTHLY_LIMIT constant.
// Otherwise, their posts limit will not be reset and no error will be returned.
// This function takes a pointer to a User object and a UsersRepository object as parameters, and returns any error that may occur.
func ResetLimit(u *User, usersRepository *UsersRepository) error {
	// TODO: Should we do this?
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

	// The default/initial value in the database is "null", so we cannot take a null value and try to perform calculations on it.
	// TODO: should we do this?
	if u.LastPublishAt == nil {
		return nil
	}

	// TODO: should we do this?
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

	return usersRepository.ResetLimit(u.ID)
}

// UpdateUserProfile updates the profile of the user with the given ID using the provided payload.
// It takes four parameters, a HandlerCtx, an UpdateProfileDTO payload, an authenticatedUserID of type toolkitEntities.ID,
// and a UsersRepository, and returns an error if the update is unsuccessful.
func UpdateUserProfile(handlerCtx *configs.HandlersCtx, payload *UpdateProfileDTO, authenticatedUserID toolkitEntities.ID, usersRepository *UsersRepository) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	u := usersRepository.FindUserByID(authenticatedUserID)

	if err := UserExists(u); err != nil {
		return err
	}

	// if new value equals to prev value, do not update
	if payload.Email != u.Email {
		if err := IsEmailInUse(usersRepository.IsEmailInUse(payload.Email)); err != nil {
			return err
		}
	}

	// if new value equals to prev value, do not update
	if payload.Nick != u.Nick {
		if err := IsNickInUse(usersRepository.IsNickInUse(payload.Nick)); err != nil {
			return err
		}
	}

	if err := usersRepository.UpdateProfile(authenticatedUserID, payload); err != nil {
		return err
	}

	return nil
}

// DecodeUserToken decodes an user JWT token and returns user's ID.
func DecodeUserToken(cfg *configs.HandlersCtx) toolkitEntities.DecodeUserTokenResult {
	claims := jwt.MapClaims{}

	t := strings.Split(cfg.C.Get("Authorization"), "Bearer ")

	if len(t) == 1 {
		return toolkitEntities.DecodeUserTokenResult{}
	}

	jwt.ParseWithClaims(t[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return cfg.Cfg.JWT.Secret, nil
	})

	parsedID, _ := toolkitEntities.ParseID(claims["id"].(string))

	u := toolkitEntities.DecodeUserTokenResult{
		ID: parsedID,
	}

	return u
}

// GetUserByToken decodes a token and get user info from token.
func GetUserByToken(cfg *configs.HandlersCtx) toolkitEntities.DecodeUserTokenResult {
	return DecodeUserToken(cfg)
}
