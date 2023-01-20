package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Questions is a model for each question in app.
type Question struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Content string             `json:"content"`

	SendTo string `json:"sendTo"`
	SentBy string `json:"sentBy,omitempty"`
	Reply  Reply  `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}
