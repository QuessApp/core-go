package repositories

import (
	"context"
	collections "core/internal/constants"
	appEntities "core/internal/entities"
	pkgEntities "core/pkg/entities"

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
func (a Auth) SignUp(payload appEntities.User) error {
	coll := a.db.Collection(collections.USERS)

	user := appEntities.User{
		ID:              pkgEntities.NewID(),
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
	coll := a.db.Collection(collections.USERS)

	var user appEntities.User

	coll.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	return user.Email != ""
}
