package auth

import (
	"context"
	"core/internal/users"

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
