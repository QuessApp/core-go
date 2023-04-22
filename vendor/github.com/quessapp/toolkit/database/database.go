package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect handles database connection.
// It returns a pointer to the database and an error.
// If the error is not nil, the database pointer is nil.
// Accepts two params: URI and DBName.
// URI is the connection string to the database.
// DBName is the name of the database.
func Connect(URI, DBName string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))

	if err != nil {
		return nil, err
	}

	return client.Database(DBName), nil
}
