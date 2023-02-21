package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// ID is an alias for ObjectID.
type ID = primitive.ObjectID

// NewID returns new ObjectID.
func NewID() ID {
	return primitive.NewObjectID()
}

// ParseID parses id from string to ObjectID.
func ParseID(id string) (ID, error) {
	return primitive.ObjectIDFromHex(id)
}
