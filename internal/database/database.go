package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect handles database connection.
func Connect(URI, DBName string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))

	if err != nil {
		return nil, err
	}

	return client.Database(DBName), nil
}
