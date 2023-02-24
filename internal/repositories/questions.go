package repositories

import (
	"context"
	collections "core/internal/constants"
	"core/internal/dtos"

	internalEntities "core/internal/entities"
	pkgEntities "core/pkg/entities"

	"time"

	"go.mongodb.org/mongo-driver/bson"
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
func (q Questions) Create(payload *dtos.CreateQuestionDTO) error {
	coll := q.db.Collection(collections.QUESTIONS)

	payload.ID = pkgEntities.NewID()
	payload.CreatedAt = time.Now()

	question := internalEntities.Question{
		ID:          payload.ID,
		Content:     payload.Content,
		IsAnonymous: payload.IsAnonymous,
		SendTo:      payload.SendTo,
		CreatedAt:   payload.CreatedAt,
		SentBy:      payload.SentBy,
		Reply:       nil,
	}

	_, err := coll.InsertOne(context.TODO(), question)

	return err
}

// FindByID finds a question by id.
func (q Questions) FindByID(id pkgEntities.ID) *internalEntities.Question {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: id}}

	question := internalEntities.Question{}

	coll.FindOne(context.Background(), filter).Decode(&question)

	return &question
}
