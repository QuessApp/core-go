package entities

import (
	toolkitEntities "github.com/kuriozapp/toolkit/entities"
)

// Report is a model for each report in app.
type Report struct {
	ID          toolkitEntities.ID `json:"id" bson:"_id"`
	Reason      string             `json:"reason"`
	Label       string             `json:"label"`
	Description string             `json:"description"`
	SentBy      User               `json:"sentBy"`
}
