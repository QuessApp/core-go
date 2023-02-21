package repositories

import (
	"context"
	"core/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Users represents users repository.
type Users struct {
	db *mongo.Database
}

// NewUsersRepository returns users repository.
func NewUsersRepository(db *mongo.Database) *Users {
	return &Users{db}
}

// FindUserByEmail finds an user by their email.
func (u Users) FindUserByEmail(email string) *entities.User {
	coll := u.db.Collection("users")

	var foundUser entities.User

	coll.FindOne(context.Background(),
		bson.M{
			"email": bson.M{"$eq": email},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByNick finds an user by their nick.
func (u Users) FindUserByNick(nick string) *entities.User {
	coll := u.db.Collection("users")

	var foundUser entities.User

	coll.FindOne(context.Background(),
		bson.M{
			"nick": bson.M{"$eq": nick},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByID finds an user by id.
func (u Users) FindUserByID(id string) *entities.User {
	coll := u.db.Collection("users")

	var foundUser entities.User

	coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&foundUser)

	return &foundUser
}

// IsNickInUse checks is an user already take a nick.
func (u Users) IsNickInUse(nick string) bool {
	coll := u.db.Collection("users")

	var user entities.User

	coll.FindOne(context.Background(), bson.D{{Key: "nick", Value: nick}}).Decode(&user)

	return user.Nick != ""
}
