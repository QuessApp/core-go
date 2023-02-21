package repositories

import (
	"context"
	"core/internal/entities"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// RegisterUser registers a new user in database.
func (a Auth) RegisterUser(payload entities.User) error {
	coll := a.db.Collection("users")

	user := entities.User{
		ID:              primitive.NewObjectID(),
		Nick:            payload.Nick,
		Name:            payload.Name,
		Email:           payload.Email,
		Password:        payload.Password,
		PostsLimit:      30,
		EnableAppEmails: true,
		IsShadowBanned:  false,
		IsPro:           false,
		CreatedAt:       time.Now(),
		AvatarURL:       "",
		CustomerID:      nil,
		LastPublishAt:   nil,
		SubscriptionID:  nil,
		ProExpiresAt:    nil,
	}

	_, err := coll.InsertOne(context.Background(), user)

	return err
}

// IsEmailInUse checks is an user already take an email.
func (a Auth) IsEmailInUse(email string) bool {
	coll := a.db.Collection("users")

	var user entities.User

	coll.FindOne(context.Background(), bson.D{{"email", email}}).Decode(&user)

	return user.Email != ""
}
