package entities

import (
	"core/pkg/entities"
	"time"
)

// Questions is a model for each question in app.
type Question struct {
	ID      entities.ID `json:"id" bson:"_id"`
	Content string      `json:"content"`

	SendTo *User `json:"sendTo" bson:"sendTo"`
	SentBy *User `json:"sentBy,omitempty" bson:"sentBy"`
	Reply  Reply `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous" bson:"IsAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty" bson:"isHiddenByReceiver"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
