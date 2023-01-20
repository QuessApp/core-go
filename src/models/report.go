package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Report is a model for each report in our app
type Report struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Reason      string             `json:"reason"`
	Label       string             `json:"label"`
	Description string             `json:"description"`
	SentBy      User               `json:"sentBy"`
}
