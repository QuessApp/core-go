package repositories

import (
	"context"
	collections "core/internal/constants"
	internal "core/internal/entities"
	pkg "core/pkg/entities"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Questions represents questions repository.
type Questions struct {
	db *mongo.Database
}

// NewQuestionsRepository returns questions repository.
func NewQuestionsRepository(db *mongo.Database) *Questions {
	return &Questions{db}
}

// Create creates a question in database.
func (q Questions) Create(payload *internal.Question) error {
	questionsColl := q.db.Collection(collections.QUESTIONS)

	payload.ID = pkg.NewID()
	payload.CreatedAt = time.Now()

	question := internal.Question{
		ID:          payload.ID,
		Content:     payload.Content,
		IsAnonymous: payload.IsAnonymous,
		SendTo:      payload.SendTo,
		CreatedAt:   payload.CreatedAt,
		SentBy:      payload.SentBy,

		Reply: nil,
	}

	_, err := questionsColl.InsertOne(context.TODO(), question)

	return err
}
