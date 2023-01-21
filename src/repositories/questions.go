package repositories

import (
	"context"
	"core/src/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type questions struct {
	db *mongo.Database
}

func NewQuestionsRepository(db *mongo.Database) *questions {
	return &questions{db}
}

// Create creates a question in database.
func (q questions) Create(payload models.Question) (*mongo.InsertOneResult, error) {
	questionsColl := q.db.Collection("questions")

	question := models.Question{
		Content:     payload.Content,
		IsAnonymous: payload.IsAnonymous,
		SendTo:      payload.SendTo,
		CreatedAt:   time.Now(),
		SentBy:      "caioamfr",
	}

	result, err := questionsColl.InsertOne(context.TODO(), question)

	return result, err
}
