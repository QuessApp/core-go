package repositories

import (
	"context"
	"core/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// auth represents auth repository.
type auth struct {
	db *mongo.Database
}

// NewAuthRepository creates auth repository.
func NewAuthRepository(db *mongo.Database) *auth {
	return &auth{db}
}

// RegisterUser registers a new user in database.
func (a auth) RegisterUser(payload models.User) (*mongo.InsertOneResult, error) {
	coll := a.db.Collection("users")

	user := models.User{
		ID:              primitive.NewObjectID(),
		Nick:            payload.Nick,
		Name:            payload.Name,
		Email:           payload.Email,
		PostsLimit:      30,
		EnableAppEmails: true,
		IsShadowBanned:  false,
		IsPro:           false,
		CreatedAt:       time.Now(),
		AvatarUrl:       "",
	}

	return coll.InsertOne(context.TODO(), user)
}