package auth

import (
	"context"

	"github.com/google/uuid"
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
		TrustedIPs:      []string{},
	}

	_, err := coll.InsertOne(context.Background(), user)

	return &user, err
}

// UpdateUserPassword updates the password for a user with the given userID.
// It takes in the userID and the newHashedPassword as parameters and updates the password
// of the user in the database with the newHashedPassword.
// It returns an error if the update fails.
func (a AuthRepository) UpdateUserPassword(userID toolkitEntities.ID, newHashedPassword []byte) error {
	coll := a.db.Collection(toolkitConstants.USERS)

	update := bson.D{
		{
			Key:   "$set",
			Value: bson.D{{Key: "password", Value: string(newHashedPassword)}},
		},
	}

	_, err := coll.UpdateByID(context.Background(), userID, update)

	return err
}

// CreateUserToken function creates a new JWT token with a given user ID, expiration time and secret key and returns it as a signed string.
// It takes a user ID, an expiration time and a secret string as arguments.
// The function first creates a MapClaims object with "id" and "exp" fields set to the provided user ID and expiration time, respectively.
// Next, it creates a new JWT token using jwt.NewWithClaims method, with the HS256 signing method and the claims object.
// Finally, the function signs the token with the provided secret key using the token.SignedString method and returns the signed token as a string.
// If any error occurs during the token creation or signing process, the function returns an empty string and the error.
func (a *AuthRepository) CreateUserToken(userID toolkitEntities.ID, expiresIn time.Time, secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": expiresIn.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// CreateAccessToken function generates a new access token for a given user and returns it as a string.
// It takes a user ID and a secret string as arguments.
// The function calls the CreateUserToken method of the AuthRepository with the given user ID, an expiration time 1 day in the future,
// and the secret string.
// If the CreateUserToken function returns an error, the function returns an empty string and the error.
// Otherwise, it returns the generated access token as a string.
func (a *AuthRepository) CreateAccessToken(userID toolkitEntities.ID, secret string) (string, error) {
	return a.CreateUserToken(userID, time.Now().Add(toolkitConstants.ONE_DAY_IN_HOURS), secret)
}

// CreateRefreshToken function generates a new refresh token for a given user and returns it as a string.
// It takes a user ID and a secret string as arguments.
// The function calls the CreateUserToken method of the AuthRepository with the given user ID, an expiration time 30 days in the future,
// and the secret string.
// If the CreateUserToken function returns an error, the function returns an empty string and the error.
// Otherwise, it returns the generated refresh token as a string.
func (a *AuthRepository) CreateRefreshToken(userID toolkitEntities.ID, secret string) (string, error) {
	return a.CreateUserToken(userID, time.Now().Add(toolkitConstants.THIRTY_DAYS_IN_HOURS), secret)
}

// CreateCodeToken creates a code token with followed fields:
// id, type, expiresAt, createdAt, createdBy and code. It returns Token and error.
// It also inserts the token into the database.
// The code is a random string generated by uuid.
func (a *AuthRepository) CreateCodeToken(userID toolkitEntities.ID) (*Token, error) {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	code := Token{
		ID:        toolkitEntities.NewID(),
		Type:      "Code",
		ExpiresAt: time.Now().Add(time.Minute * 10),
		CreatedAt: time.Now(),
		CreatedBy: &userID,
		Code:      uuid.New().String()[0:5],
	}

	_, err := coll.InsertOne(context.Background(), code)

	if err != nil {
		return nil, err
	}

	return &code, nil
}

// CreateAuthTokens function creates a new token pair (access token and refresh token) and saves them in the database.
// It takes a user ID and a secret string as arguments.
// The function first creates an access token using the CreateAccessToken function of the AuthRepository.
// Then, it creates a refresh token using the CreateRefreshToken function of the AuthRepository.
// Next, it creates a Token object with the generated tokens, expiration date, creation date, user ID, and type ("Bearer").
// It then inserts the token object into the tokens collection of the database using MongoDB driver's InsertOne method.
// If the insertion is successful, the function sets the access token in the token object and returns it.
// If any error occurs, the function returns nil and the error.
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

// DeleteRefreshToken deletes a refresh token from the database.
// It takes in the refresh token as a parameter and returns an error if one occurs.
func (a AuthRepository) DeleteRefreshToken(token string) error {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "refreshToken", Value: token,
		},
	}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// CheckIfTrustedIPExists checks if a given IP address exists in the user's trusted IPs list.
// It takes in the user ID and the IP address as parameters and returns true if the IP address exists in the list.
// Otherwise, it returns false.
func (a AuthRepository) CheckIfTrustedIPExists(userID toolkitEntities.ID, ip string) bool {
	coll := a.db.Collection(toolkitConstants.USERS)

	filter := bson.D{
		{
			Key: "id", Value: userID,
		},
		{
			Key: "trustedIps", Value: ip,
		},
	}

	count, _ := coll.CountDocuments(context.Background(), filter)

	return count > 0
}

// AddNewTrustedIPIfDontExists adds a new trusted IP to the user's trusted IPs list if it does not already exist.
// It takes in the user ID and the IP address as parameters and returns an error if one occurs.
func (a AuthRepository) AddNewTrustedIPIfDontExists(userID toolkitEntities.ID, ip string) error {
	coll := a.db.Collection(toolkitConstants.USERS)

	filter := bson.D{
		{
			Key: "_id", Value: userID,
		},
	}

	update := bson.D{
		{
			Key: "$addToSet", Value: bson.D{
				{Key: "trustedIps", Value: bson.D{
					{Key: "$each", Value: []string{ip}}},
				},
			},
		},
	}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// FindTokenByCode finds a token in the database that matches the given code.
// It takes in the code as a parameter and returns a pointer to a Token object if it exists,
// or nil if it does not exist.
func (a AuthRepository) FindTokenByCode(code string) *Token {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "type", Value: "Code",
		},
		{
			Key:   "code",
			Value: code,
		},
	}

	t := Token{}

	coll.FindOne(context.Background(), filter).Decode(&t)

	return &t
}

// DeleteByID removes a token from the database collection "tokens" that matches the given ID.
// It returns an error if there was a problem deleting the token, or nil if the token was deleted successfully.
func (a AuthRepository) DeleteTokenByID(ID toolkitEntities.ID) error {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	filter := bson.D{
		{
			Key: "_id", Value: ID,
		},
	}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// DeleteAllUserTokens removes all specified type tokens from the database collection "tokens" that match the given userID.
// It returns an error if there was a problem deleting the tokens, or nil if the tokens were deleted successfully.
func (a AuthRepository) DeleteAllUserTokens(userID toolkitEntities.ID, tokenType *string) error {
	coll := a.db.Collection(toolkitConstants.TOKENS)

	if tokenType == nil {
		*tokenType = "Bearer"
	}

	filter := bson.D{
		{
			Key: "createdBy", Value: userID,
		},
		{
			Key: "type", Value: &tokenType,
		},
	}

	_, err := coll.DeleteMany(context.Background(), filter)

	return err
}
