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
	"go.mongodb.org/mongo-driver/mongo/options"
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

// FindByID finds a question by id in database.
func (q Questions) FindByID(id pkgEntities.ID) *internalEntities.Question {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: id}}

	question := internalEntities.Question{}

	coll.FindOne(context.Background(), filter).Decode(&question)

	return &question
}

// GetAll gets all paginated questions from database.
func (q Questions) GetAll(page *int64, sort, filter *string, authenticatedUserId pkgEntities.ID) (*internalEntities.PaginatedQuestions, error) {
	var LIMIT int64 = 30

	coll := q.db.Collection(collections.QUESTIONS)

	findFilterOptions := bson.D{{Key: "sendTo", Value: authenticatedUserId}, {Key: "isReplied", Value: false}}

	if *filter == "sent" {
		findFilterOptions = bson.D{{Key: "sentBy", Value: authenticatedUserId}, {Key: "isReplied", Value: false}}
	}

	if *filter == "replied" {
		findFilterOptions = bson.D{{Key: "isReplied", Value: false}}
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})

	if *sort == "desc" {
		findOptions = options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	}

	findOptions.SetSkip((*page - 1) * LIMIT)
	findOptions.SetLimit(LIMIT)

	questions := []internalEntities.Question{}

	cursor, err := coll.Find(context.Background(), findFilterOptions, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &questions); err != nil {
		return nil, err
	}

	totalCount, err := coll.CountDocuments(context.Background(), findFilterOptions)

	if err != nil {
		return nil, err
	}

	result := internalEntities.PaginatedQuestions{
		TotalCount: totalCount,
		Questions:  &questions,
	}

	return &result, nil
}

// Delete deletes a question from database.
func (q Questions) Delete(id pkgEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}
