package repositories

import (
	"context"
	"core/internal/entities"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type questions struct {
	db *mongo.Database
}

// NewQuestionsRepository returns questions repository.
func NewQuestionsRepository(db *mongo.Database) *questions {
	return &questions{db}
}

// Create creates a question in database.
func (q questions) Create(payload entities.Question) (*mongo.InsertOneResult, error) {
	questionsColl := q.db.Collection("questions")

	question := entities.Question{
		Content:     payload.Content,
		IsAnonymous: payload.IsAnonymous,
		SendTo:      payload.SendTo,
		CreatedAt:   time.Now(),
		SentBy:      payload.SentBy,
	}

	result, err := questionsColl.InsertOne(context.TODO(), question)

	return result, err
}
