package questions

import (
	"context"

	collections "github.com/kuriozapp/toolkit/constants"

	toolkitEntities "github.com/kuriozapp/toolkit/entities"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionsRepository represents questions repository.
type QuestionsRepository struct {
	db *mongo.Database
}

// NewRepository returns questions repository.
func NewRepository(db *mongo.Database) *QuestionsRepository {
	return &QuestionsRepository{db}
}

// Create creates a question in database.
func (q QuestionsRepository) Create(payload *CreateQuestionDTO) error {
	coll := q.db.Collection(collections.QUESTIONS)

	payload.ID = toolkitEntities.NewID()
	payload.CreatedAt = time.Now()

	question := Question{
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

// FindQuestionByID finds a question by id in database.
func (q QuestionsRepository) FindQuestionByID(id toolkitEntities.ID) *Question {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: id}}

	question := Question{}

	coll.FindOne(context.Background(), filter).Decode(&question)

	return &question
}

// GetAll gets all paginated questions from database.
func (q QuestionsRepository) GetAll(page *int64, sort, filter *string, authenticatedUserId toolkitEntities.ID) (*PaginatedQuestions, error) {
	var LIMIT int64 = 30

	coll := q.db.Collection(collections.QUESTIONS)

	findFilterOptions := bson.D{
		{Key: "sendTo", Value: authenticatedUserId},
		{Key: "isReplied", Value: false},
		{Key: "isHiddenByReceiver", Value: false},
	}

	if *filter == "sent" {
		findFilterOptions = bson.D{
			{Key: "sentBy", Value: authenticatedUserId},
			{Key: "isReplied", Value: false},
			{Key: "isHiddenByReceiver", Value: false},
		}
	}

	if *filter == "replied" {
		findFilterOptions = bson.D{
			{Key: "isReplied", Value: true},
			{Key: "isHiddenByReceiver", Value: false},
		}
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})

	if *sort == "desc" {
		findOptions = options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	}

	findOptions.SetSkip((*page - 1) * LIMIT)
	findOptions.SetLimit(LIMIT)

	countOptions := options.Count()

	countOptions.SetSkip((*page - 1) * LIMIT)
	countOptions.SetLimit(LIMIT)

	questions := []Question{}

	cursor, err := coll.Find(context.Background(), findFilterOptions, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &questions); err != nil {
		return nil, err
	}

	totalCount, err := coll.CountDocuments(context.Background(), findFilterOptions, countOptions)

	if err != nil {
		return nil, err
	}

	result := PaginatedQuestions{
		TotalCount: totalCount,
		Questions:  &questions,
	}

	return &result, nil
}

// Delete deletes a question from database.
func (q QuestionsRepository) Delete(id toolkitEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// Hide hides a question.
func (q QuestionsRepository) Hide(id toolkitEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isHiddenByReceiver", Value: true}}}}

	_, err := coll.UpdateByID(context.Background(), id, update)

	return err
}

// Reply replies a question.
func (q QuestionsRepository) Reply(payload *ReplyQuestionDTO) error {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: payload.ID}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "isReplied", Value: true},
				{Key: "reply", Value: payload.Content},
				{Key: "repliedAt", Value: time.Now()},
			},
		},
	}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
