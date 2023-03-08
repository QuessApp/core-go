package users

import (
	"context"
	"time"

	collections "github.com/kuriozapp/toolkit/constants"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Search searchs for an user by their nick or name.
func (u UsersRepository) Search(value string, page *int64) (*PaginatedUsers, error) {
	if value == "" {
		return &PaginatedUsers{
			Users: &[]User{},
		}, nil
	}

	var LIMIT int64 = 30

	coll := u.db.Collection(collections.USERS)

	findFilterOptions := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "name", Value: primitive.Regex{Pattern: value, Options: ""}}},
			bson.D{{Key: "nick", Value: primitive.Regex{Pattern: value, Options: ""}}},
		}},
	}

	findOptions := options.Find().SetSort(bson.D{
		{Key: "nick", Value: 1},
		{Key: "name", Value: 1},
	}).SetProjection(bson.D{
		{Key: "id", Value: 1},
		{Key: "nick", Value: 1},
		{Key: "avatarUrl", Value: 1},
		{Key: "name", Value: 1},
	})

	findOptions.SetSkip((*page - 1) * LIMIT)
	findOptions.SetLimit(LIMIT)

	countOptions := options.Count()

	countOptions.SetSkip((*page - 1) * LIMIT)
	countOptions.SetLimit(LIMIT)

	users := []User{}

	cursor, err := coll.Find(context.Background(), findFilterOptions, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	totalCount, err := coll.CountDocuments(context.Background(), findFilterOptions, countOptions)

	if err != nil {
		return nil, err
	}

	result := PaginatedUsers{
		TotalCount: totalCount,
		Users:      &users,
	}

	return &result, nil
}

// DecrementLimit decrements user's post limit if user is not a PRO member.
func (u *UsersRepository) DecrementLimit(userId toolkitEntities.ID, newValue int) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "postsLimit", Value: newValue}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// UpdatePreferences updates user preferences such as emails, etc.
func (u *UsersRepository) UpdatePreferences(userId toolkitEntities.ID, payload *UpdatePreferencesDTO) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{
			Key: "enableAppEmails", Value: payload.EnableAPPEmails,
		},
		{
			Key: "enableAppPushNotifications", Value: payload.EnanbleAPPPushNotifications,
		},
	}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// UpdateLastPublishedAt updates last publish at field in database.
func (u *UsersRepository) UpdateLastPublishedAt(userId toolkitEntities.ID) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{
			Key: "lastPublishAt", Value: time.Now(),
		},
	}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
