package repositories

import (
	"context"
	"core/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// users represents users repository.
type users struct {
	db *mongo.Database
}

// NewUsersRepository creates users repository.
func NewUsersRepository(db *mongo.Database) *users {
	return &users{db}
}

// Create creates a new user in database.
func (u users) Create(payload models.User) (*mongo.InsertOneResult, error) {
	coll := u.db.Collection("users")

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
