package entities

import "core/pkg/entities"

// Reply is a model for each reply in app.
type Reply struct {
	ID        entities.ID `json:"id" bson:"_id"`
	Content   string      `json:"content"`
	Replied   bool        `json:"replied"`
	RepliedBy string      `json:"repliedBy"`
}
