package users

import (
	"context"

	collections "github.com/kuriozapp/toolkit/constants"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UsersRepository represents users repository.
type UsersRepository struct {
	db *mongo.Database
}

// NewRepository returns users repository.
func NewRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{db}
}

// FindUserByEmail finds an user by their email.
func (u UsersRepository) FindUserByEmail(email string) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(),
		bson.M{
			"email": bson.M{"$eq": email},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByNick finds an user by their nick.
func (u UsersRepository) FindUserByNick(nick string) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(),
		bson.M{
			"nick": bson.M{"$eq": nick},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByID finds an user by id.
func (u UsersRepository) FindUserByID(userId toolkitEntities.ID) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&foundUser)

	return &foundUser
}

// IsNickInUse checks is an user already take a nick.
func (u UsersRepository) IsNickInUse(nick string) bool {
	coll := u.db.Collection(collections.USERS)

	var user User

	coll.FindOne(context.Background(), bson.D{{Key: "nick", Value: nick}}).Decode(&user)

	return user.Nick != ""
}

// DecrementLimit decrements user's post limit if user is not a PRO member.
func (u *UsersRepository) DecrementLimit(userId toolkitEntities.ID, newValue int) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "postsLimit", Value: newValue}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
