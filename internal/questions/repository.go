package questions

import (
	"context"

	collections "github.com/quessapp/toolkit/constants"

	toolkitEntities "github.com/quessapp/toolkit/entities"

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

// Create creates a new question in the database with the given payload.
// It returns an error if the insertion operation fails.
func (q QuestionsRepository) Create(payload *CreateQuestionDTO) error {
	coll := q.db.Collection(collections.QUESTIONS)

	payload.ID = toolkitEntities.NewID()
	payload.CreatedAt = time.Now()
	var repliedAt *time.Time

	question := Question{
		ID:             payload.ID,
		Content:        payload.Content,
		IsAnonymous:    payload.IsAnonymous,
		SendTo:         payload.SendTo,
		CreatedAt:      payload.CreatedAt,
		SentBy:         payload.SentBy,
		Reply:          nil,
		RepliedAt:      repliedAt,
		RepliesHistory: []ReplyHistory{},
	}

	_, err := coll.InsertOne(context.Background(), question)

	return err
}

// FindQuestionByID finds a question in the database by its ID.
// It returns a pointer to the Question found, or nil if no question was found.
func (q QuestionsRepository) FindQuestionByID(ID toolkitEntities.ID) *Question {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: ID}}

	question := Question{}

	coll.FindOne(context.Background(), filter).Decode(&question)

	return &question
}

// GetAll returns a paginated list of questions from the questions collection. It takes
// a page number (int64), a sort (string), a filter (string) and an authenticatedUserID (toolkitEntities.ID)
// as arguments and returns a pointer to a PaginatedQuestions struct and an error. The function retrieves
// the corresponding documents from the questions collection based on the filter (sent, replied or all),
// the sort (asc or desc) and the page number, and returns them as a list of Question structs, along with
// the total number of documents found in the collection that match the given filter. The function also
// returns an error if the database query fails.
func (q QuestionsRepository) GetAll(page *int64, sort, filter *string, authenticatedUserID toolkitEntities.ID) (*PaginatedQuestions, error) {
	var LIMIT int64 = 30

	coll := q.db.Collection(collections.QUESTIONS)

	findFilterOptions := bson.D{
		{Key: "sendTo", Value: authenticatedUserID},
		{Key: "isReplied", Value: false},
		{Key: "isHiddenByReceiver", Value: false},
	}

	if *filter == "sent" {
		findFilterOptions = bson.D{
			{Key: "sentBy", Value: authenticatedUserID},
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
	findOptions.SetProjection(bson.D{{Key: "repliesHistory", Value: 0}})

	countOptions := options.Count()

	countOptions.SetSkip((*page - 1) * LIMIT)
	countOptions.SetLimit(LIMIT)

	questions := []Question{}

	cursor, err := coll.Find(context.Background(), findFilterOptions, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &questions); err != nil {
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

// Delete deletes a question from the questions collection. It takes a
// toolkitEntities.ID as argument and returns an error. The function deletes the
// corresponding document in the questions collection and returns an error if the
// delete operation fails.
func (q QuestionsRepository) Delete(ID toolkitEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: ID}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

// Hide hides a question from the receiver's feed by setting the isHiddenByReceiver
// field to true. It takes a toolkitEntities.ID as argument and returns an error.
// The function updates the corresponding document in the questions collection
// and returns an error if the update operation fails.
func (q QuestionsRepository) Hide(ID toolkitEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isHiddenByReceiver", Value: true}}}}

	_, err := coll.UpdateByID(context.Background(), ID, update)

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

// EditReply updates the content of a reply to a question and adds the old content
// to the repliesHistory field. It takes a pointer to an EditQuestionReplyDTO as
// argument and returns an error. The function first creates a ReplyHistory slice
// containing the old and new contents, and then updates the reply and
// repliesHistory fields in the corresponding document in the questions collection.
// The function returns an error if the update operation fails.
func (q QuestionsRepository) EditReply(payload *EditQuestionReplyDTO) error {
	coll := q.db.Collection(collections.QUESTIONS)

	addHistory := []ReplyHistory{
		// create history for old content
		{
			ID:        toolkitEntities.NewID(),
			CreatedAt: payload.OldContentCreatedAt,
			Content:   payload.OldContent,
		},
		{
			ID:        toolkitEntities.NewID(),
			CreatedAt: time.Now(),
			Content:   payload.Content,
		},
	}

	filter := bson.D{{Key: "_id", Value: payload.ID}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "reply", Value: payload.Content},
			},
		},
		{
			Key: "$push", Value: bson.D{
				{Key: "repliesHistory", Value: bson.D{
					{Key: "$each", Value: addHistory},
				}},
			},
		},
	}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}

// RemoveReply removes the reply to a question with the given ID from the Questions collection.
// It requires a toolkitEntities.ID object as input parameter.
// It returns an error if the reply cannot be removed from the collection.
func (q QuestionsRepository) RemoveReply(ID toolkitEntities.ID) error {
	coll := q.db.Collection(collections.QUESTIONS)

	filter := bson.D{{Key: "_id", Value: ID}}

	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "reply", Value: nil},
				{Key: "isReplied", Value: false},
				{Key: "repliedAt", Value: nil},
				{Key: "repliesHistory", Value: []ReplyHistory{}},
			},
		},
	}

	_, err := coll.UpdateOne(context.Background(), filter, update)

	return err
}
