package repositories

import (
	"context"
	collections "core/internal/constants"
	internal "core/internal/entities"
	pkg "core/pkg/entities"

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
func (u Users) FindUserByEmail(email string) *internal.User {
	coll := u.db.Collection(collections.USERS)

	var foundUser internal.User

	coll.FindOne(context.Background(),
		bson.M{
			"email": bson.M{"$eq": email},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByNick finds an user by their nick.
func (u Users) FindUserByNick(nick string) *internal.User {
	coll := u.db.Collection(collections.USERS)

	var foundUser internal.User

	coll.FindOne(context.Background(),
		bson.M{
			"nick": bson.M{"$eq": nick},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByID finds an user by id.
func (u Users) FindUserByID(userId pkg.ID) (*internal.User, error) {
	coll := u.db.Collection(collections.USERS)

	var foundUser internal.User

	err := coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: userId}}).Decode(&foundUser)

	return &foundUser, err
}

// IsNickInUse checks is an user already take a nick.
func (u Users) IsNickInUse(nick string) bool {
	coll := u.db.Collection(collections.USERS)

	var user internal.User

	coll.FindOne(context.Background(), bson.D{{Key: "nick", Value: nick}}).Decode(&user)

	return user.Nick != ""
}

// DecrementLimit decrements user's post limit if user is not a PRO member.
func (u *Users) DecrementLimit(userId pkg.ID, newValue int) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "postsLimit", Value: newValue}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
