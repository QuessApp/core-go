package entities

import (
	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// Reply is a model for each reply in app.
type Reply struct {
	ID        toolkitEntities.ID `json:"id" bson:"_id"`
	Content   string             `json:"content"`
	Replied   bool               `json:"replied"`
	RepliedBy *User              `json:"repliedBy" bson:"repliedBy"`
}
