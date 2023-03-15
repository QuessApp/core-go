package auth

import (
	"context"

	"github.com/quessapp/core-go/internal/users"
	collections "github.com/quessapp/toolkit/constants"
	toolkitEntities "github.com/quessapp/toolkit/entities"

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
	coll := a.db.Collection(collections.USERS)

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
