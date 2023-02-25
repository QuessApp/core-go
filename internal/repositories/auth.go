package repositories

import (
	"context"
	collections "core/internal/constants"
	"core/internal/dtos"
	internalEntities "core/internal/entities"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Auth represents auth repository.
type Auth struct {
	db *mongo.Database
}

// NewAuthRepository returns auth repository.
func NewAuthRepository(db *mongo.Database) *Auth {
	return &Auth{db}
}

// SignUp registers a new user in database.
func (a Auth) SignUp(payload *dtos.SignUpUserDTO) (*internalEntities.User, error) {
	coll := a.db.Collection(collections.USERS)

	payload.ID = toolkitEntities.NewID()
	payload.CreatedAt = time.Now()

	user := internalEntities.User{
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
func (a Auth) IsEmailInUse(email string) bool {
	coll := a.db.Collection(collections.USERS)

	user := internalEntities.User{}

	coll.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	return user.Email != ""
}
