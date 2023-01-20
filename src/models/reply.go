package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Reply is a model for each reply in app.
type Reply struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Content   string             `json:"content"`
	Replied   bool               `json:"replied"`
	RepliedBy User               `json:"repliedBy"`
}
