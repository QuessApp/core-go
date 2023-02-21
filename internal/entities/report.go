package entities

import (
	"core/pkg/entities"
)

// Report is a model for each report in app.
type Report struct {
	ID          entities.ID `json:"id" bson:"_id"`
	Reason      string      `json:"reason"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	SentBy      User        `json:"sentBy"`
}
