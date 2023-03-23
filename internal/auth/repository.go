package auth

import (
	"context"

	"github.com/quessapp/core-go/internal/users"
	"go.mongodb.org/mongo-driver/bson"

	toolkitConstants "github.com/quessapp/toolkit/constants"
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"github.com/golang-jwt/jwt/v4"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// AuthRepository represents auth repository.
type AuthRepository struct {
	db *mongo.Database
}

// NewAuthRepository returns auth repository.
func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db}
}

// CreateUserToken creates an user JWT token with followed fields:
// id and exp. It returns string and error.
func (a *AuthRepository) CreateUserToken(userID toolkitEntities.ID, expiresIn time.Time, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": expiresIn.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (a *AuthRepository) CreateAccessToken(userID toolkitEntities.ID, secret string) (string, error) {
	return a.CreateUserToken(userID, time.Now().Add(toolkitConstants.ONE_DAY_IN_HOURS), secret)
}

func (a *AuthRepository) CreateRefreshToken(userID toolkitEntities.ID, secret string) (string, error) {
	return a.CreateUserToken(userID, time.Now().Add(toolkitConstants.THIRTY_DAYS_IN_HOURS), secret)
}

func (a *AuthRepository) CreateAuthTokens(userID toolkitEntities.ID, secret string) (*Token, error) {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	accessToken, err := a.CreateAccessToken(userID, secret)

	if err != nil {
		return nil, err
	}

	refreshToken, err := a.CreateRefreshToken(userID, secret)

	if err != nil {
		return nil, err
	}

	tokens := Token{
		ID:           toolkitEntities.NewID(),
		Type:         "Bearer",
		ExpiresAt:    time.Now().Add(toolkitConstants.THIRTY_DAYS_IN_HOURS),
		CreatedAt:    time.Now(),
		CreatedBy:    &userID,
		RefreshToken: refreshToken,
	}

	_, err = coll.InsertOne(context.Background(), tokens)

	if err != nil {
		return nil, err
	}

	tokens.AccessToken = accessToken

	return &tokens, nil
}

// SignUp is a method of AuthRepository that creates a new user in the database.
// The method receives a SignUpUserDTO as input, which contains the user's information to be stored.
// The method generates a new ID for the user using the toolkitEntities.NewID method and sets the CreatedAt field to the current time.
// The method then creates a new users.User object with the provided data and default values for fields such as PostsLimit, EnableAPPEmails, IsShadowBanned, IsPRO, AvatarURL, CustomerID, LastPublishAt, SubscriptionID and ProExpiresAt.
// Finally, the method inserts the user in the database using the InsertOne method from the mongo-go-driver library and returns the inserted user and any error that may have occurred during the insertion.
func (a AuthRepository) SignUp(payload *SignUpUserDTO) (*users.User, error) {
	coll := a.db.Collection(toolkitConstants.USERS)

	payload.ID = toolkitEntities.NewID()
	payload.CreatedAt = time.Now()

	user := users.User{
		ID:              payload.ID,
		Nick:            payload.Nick,
		Name:            payload.Name,
		Email:           payload.Email,
		Password:        payload.Password,
		PostsLimit:      30,
		EnableAPPEmails: true,
		IsShadowBanned:  false,
		IsPRO:           false,
		CreatedAt:       &payload.CreatedAt,
		AvatarURL:       "",
		CustomerID:      nil,
		LastPublishAt:   nil,
		SubscriptionID:  nil,
		ProExpiresAt:    nil,
		Locale:          payload.Locale,
	}

	_, err := coll.InsertOne(context.Background(), user)

	return &user, err
}

// FindTokenByUserIDAndRefreshToken searches for a token in the database collection "tokens"
// that matches the given userID and refreshToken. It returns a pointer to the matching Token
// if one is found, or nil if no such token exists.
func (a AuthRepository) FindTokenByUserIDAndRefreshToken(userID toolkitEntities.ID, refreshToken string) *Token {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "createdBy", Value: userID,
		},
		{
			Key:   "refreshToken",
			Value: refreshToken,
		},
	}

	t := Token{}

	coll.FindOne(context.Background(), filter).Decode(&t)

	return &t
}

// DeleteByID removes a token from the database collection "tokens" that matches the given ID.
// It returns an error if there was a problem deleting the token, or nil if the token was deleted successfully.
func (a AuthRepository) DeleteByID(ID toolkitEntities.ID) error {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "_id", Value: ID,
		},
	}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// DeleteAllUserTokens removes all tokens from the database collection "tokens" that match the given userID.
// It returns an error if there was a problem deleting the tokens, or nil if the tokens were deleted successfully.
func (a AuthRepository) DeleteAllUserTokens(userID toolkitEntities.ID) error {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "createdBy", Value: userID,
		},
	}

	_, err := coll.DeleteMany(context.Background(), filter)

	return err
}
