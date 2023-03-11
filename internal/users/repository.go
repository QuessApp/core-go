package users

import (
	"context"
	"time"

	collections "github.com/quessapp/toolkit/constants"

	toolkitEntities "github.com/quessapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UsersRepository represents users repository.
type UsersRepository struct {
	db *mongo.Database
}

// NewRepository creates a new instance of the UsersRepository struct and returns a pointer to it.
// The function takes a pointer to a mongo.Database as an argument, which is used to initialize the UsersRepository's db field.
func NewRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{db}
}

// FindUserByEmail retrieves a user from the database based on their email and returns a pointer to the User object.
// It takes the user's email as a parameter and performs a database lookup to find the matching user.
// If the user is found, a pointer to the User object is returned. Otherwise, a nil pointer is returned.
func (u UsersRepository) FindUserByEmail(email string) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(),
		bson.M{
			"email": bson.M{"$eq": email},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByNick retrieves a user from the database based on their nickname and returns a pointer to the User object.
// It takes the user's nickname as a parameter and performs a database lookup to find the matching user.
// If the user is found, a pointer to the User object is returned. Otherwise, a nil pointer is returned.
func (u UsersRepository) FindUserByNick(nick string) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(),
		bson.M{
			"nick": bson.M{"$eq": nick},
		}).Decode(&foundUser)

	return &foundUser
}

// FindUserByID retrieves a user from the database based on their id and returns a pointer to the User object.
// It takes the user's id as a parameter and performs a database lookup to find the matching user.
// If the user is found, a pointer to the User object is returned. Otherwise, a nil pointer is returned.
func (u UsersRepository) FindUserByID(userID toolkitEntities.ID) *User {
	coll := u.db.Collection(collections.USERS)

	var foundUser User

	coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: userID}}).Decode(&foundUser)

	return &foundUser
}

// IsNickInUse checks if a user with the given nickname exists in the database.
// It takes the user's nickname as a parameter and performs a database lookup to find the matching user.
// If a user with the given nickname is found, it returns true. Otherwise, it returns false.
func (u UsersRepository) IsNickInUse(nick string) bool {
	coll := u.db.Collection(collections.USERS)

	var user User

	coll.FindOne(context.Background(), bson.D{{Key: "nick", Value: nick}}).Decode(&user)

	return user.Nick != ""
}

// IsEmailInUse checks is an user already take an email.
func (u UsersRepository) IsEmailInUse(email string) bool {
	coll := u.db.Collection(collections.USERS)

	user := User{}

	coll.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	return user.Email != ""
}

// Search searches for users whose names or nicks match the given value, and returns a paginated list of results.
// The page parameter is used to determine which page of the results to return.
// If the value parameter is an empty string, an empty list is returned.
// The function returns a pointer to a PaginatedUsers struct and an error.
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

// DecrementLimit updates the "postsLimit" field of the user document with the given ID to the provided value.
// It returns an error if the update operation fails.
func (u *UsersRepository) DecrementLimit(userID toolkitEntities.ID, newValue int) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "postsLimit", Value: newValue}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// UpdateAvatar updates the avatar URL for a user in the database.
func (u *UsersRepository) UpdateAvatar(userID toolkitEntities.ID, URI string) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "avatarUrl", Value: URI}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// ResetLimit updates the "postsLimit" field of the user document with the given ID to 30.
// It returns an error if the update operation fails.
func (u *UsersRepository) ResetLimit(userID toolkitEntities.ID) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "postsLimit", Value: USER_DEFAULT_POST_MONTHLY_LIMIT}}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// UpdateLastPublishedAt takes a user ID and a payload containing updated preferences for the user.
// It updates the corresponding user document in the database with the new preference values for "enableAppEmails" and "enableAppPushNotifications".
// It returns an error if the update operation fails.
func (u *UsersRepository) UpdatePreferences(userID toolkitEntities.ID, payload *UpdatePreferencesDTO) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}
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

// UpdateLastPublishedAt takes a user ID and updates the corresponding user document in the database with the new value for field "lastPublishAt".
// It returns an error if the update operation fails.
func (u *UsersRepository) UpdateLastPublishedAt(userID toolkitEntities.ID) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{
			Key: "lastPublishAt", Value: time.Now(),
		},
	}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// UpdateProfile updates the profile of the user with the given ID using the provided payload.
// It takes two parameters, a userID of type toolkitEntities.ID and a pointer to an UpdateProfileDTO payload.
// It returns an error if the update is unsuccessful.
func (u *UsersRepository) UpdateProfile(userID toolkitEntities.ID, payload *UpdateProfileDTO) error {
	coll := u.db.Collection(collections.USERS)

	filter := bson.D{{Key: "_id", Value: userID}}

	// TODO: update password
	update := bson.D{{Key: "$set", Value: bson.D{
		{
			Key: "nick", Value: payload.Nick,
		},
		{
			Key: "name", Value: payload.Name,
		},
		{
			Key: "locale", Value: payload.Locale,
		},
		{
			Key: "email", Value: payload.Email,
		},
	}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
