package reports

import (
	"context"
	"time"

	collections "github.com/quessapp/toolkit/constants"
	toolkitEntities "github.com/quessapp/toolkit/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReportsRepository represents reports repository.
type ReportsRepository struct {
	db *mongo.Database
}

// NewRepository returns reports repository.
func NewRepository(db *mongo.Database) *ReportsRepository {
	return &ReportsRepository{db}
}

// Create is a method of ReportsRepository that receives a CreateReportDTO payload and creates a new report based on that information.
// The method uses the ReportsRepository's db field to access the database collection of reports.
// It creates a new Report struct with the given payload information and generates a new ID for the report using the toolkitEntities.NewID() function.
// Finally, the method inserts the new report in the database collection and returns any error that occurred during the process.
func (r *ReportsRepository) Create(payload *CreateReportDTO) error {
	coll := r.db.Collection(collections.REPORTS)

	report := Report{
		ID:        toolkitEntities.NewID(),
		Type:      payload.Type,
		Reason:    payload.Reason,
		SendTo:    payload.SendTo,
		SentBy:    payload.SentBy,
		CreatedAt: time.Now(),
	}

	_, err := coll.InsertOne(context.Background(), report)

	return err
}

// AlreadySent is a method of ReportsRepository that receives a CreateReportDTO payload and returns a boolean indicating if a report has already been sent for the same content.
// The method uses the ReportsRepository's db field to access the database collection of reports.
// It creates a filter based on the payload's SentBy, SendTo and Reason fields.
// The method then attempts to find a report in the database collection that matches the filter criteria using the FindOne method.
// If a matching report is found, its ID is retrieved and checked using the toolkitEntities.IsZeroID function.
// If the ID is not zero, it means a report has already been sent for the same content and the method returns true.
// If the ID is zero, it means a report has not been sent for the same content and the method returns false.
func (r *ReportsRepository) AlreadySent(payload *CreateReportDTO) bool {
	coll := r.db.Collection(collections.REPORTS)

	filter := bson.D{
		{
			Key: "sentBy", Value: payload.SentBy,
		},
		{
			Key:   "sendTo",
			Value: payload.SendTo,
		},
		{
			Key:   "reason",
			Value: payload.Reason,
		},
	}

	foundRegistry := Report{}

	coll.FindOne(context.Background(), filter).Decode(&foundRegistry)

	return !toolkitEntities.IsZeroID(foundRegistry.ID)
}

// Delete removes a report with the given reportID from the database.
// It returns an error if the operation fails.
func (r *ReportsRepository) Delete(reportID toolkitEntities.ID) error {
	coll := r.db.Collection(collections.REPORTS)

	filter := bson.D{{Key: "_id", Value: reportID}}

	_, err := coll.DeleteOne(context.Background(), filter)

	return err
}

func (r *ReportsRepository) FindByID(reportID toolkitEntities.ID) (*Report, error) {
	coll := r.db.Collection(collections.REPORTS)

	filter := bson.D{{Key: "_id", Value: reportID}}

	foundRegistry := Report{}

	err := coll.FindOne(context.Background(), filter).Decode(&foundRegistry)

	return &foundRegistry, err
}