package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// ID is an alias for ObjectID.
type ID = primitive.ObjectID

// NewID returns new Object ID.
func NewID() ID {
	return primitive.NewObjectID()
}
