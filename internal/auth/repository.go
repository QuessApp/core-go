package auth

import (
	"context"
	"core/internal/users"
	collections "core/pkg/constants"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// SignUp registers a new user in database.
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
		EnableAppEmails: true,
		IsShadowBanned:  false,
		IsPRO:           false,
		CreatedAt:       &payload.CreatedAt,
		AvatarURL:       "",
		CustomerID:      nil,
		LastPublishAt:   nil,
		SubscriptionID:  nil,
		ProExpiresAt:    nil,
	}

	_, err := coll.InsertOne(context.Background(), user)

	return &user, err
}

// IsEmailInUse checks is an user already take an email.
func (a AuthRepository) IsEmailInUse(email string) bool {
	coll := a.db.Collection(collections.USERS)

	user := users.User{}

	coll.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	return user.Email != ""
}
