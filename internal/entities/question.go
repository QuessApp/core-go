package entities

import (
	"core/pkg/entities"
	"time"
)

// Questions is a model for each question in app.
type Question struct {
	ID      entities.ID `json:"id" bson:"_id"`
	Content string      `json:"content"`

	SendTo string `json:"sendTo"`
	SentBy string `json:"sentBy,omitempty"`
	Reply  Reply  `json:"reply,omitempty"`

	IsAnonymous        bool `json:"isAnonymous"`
	IsHiddenByReceiver bool `json:"isHiddenByReceiver,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}
