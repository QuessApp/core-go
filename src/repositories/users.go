package repositories

import (
	"context"
	"core/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// users represents users repository.
type Users struct {
	db *mongo.Database
}

// NewUsersRepository creates users repository.
func NewUsersRepository(db *mongo.Database) *Users {
	return &Users{db}
}

// FindUserByEmail finds an user by their email.
func (u Users) FindUserByEmail(email string) models.User {
	coll := u.db.Collection("users")

	var foundUser models.User

	coll.FindOne(context.TODO(),
		bson.M{
			"email": bson.M{"$eq": email},
		}).Decode(&foundUser)

	return foundUser
}

// FindUserByNick finds an user by their nick.
func (u Users) FindUserByNick(nick string) models.User {
	coll := u.db.Collection("users")

	var foundUser models.User

	coll.FindOne(context.TODO(),
		bson.M{
			"nick": bson.M{"$eq": nick},
		}).Decode(&foundUser)

	return foundUser
}
