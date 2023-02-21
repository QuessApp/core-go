package entities

import (
	"core/pkg/entities"
	"time"
)

// Questions is a model for each question in app.
type Question struct {
	ID      entities.ID `json:"id" bson:"_id"`
	Content string      `json:"content"`

	SendTo string `json:"sendTo" bson:"sendTo"`
	SentBy string `json:"sentBy,omitempty" bson:"sentBy"`
	// Reply is replied data content. Type must be Entities.Reply or nil.
	Reply any `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous" bson:"IsAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty" bson:"isHiddenByReceiver"`
	IsReplied          bool `json:"isReplied" bson:"isReplied"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
