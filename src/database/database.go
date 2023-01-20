package database

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mutex sync.Mutex
var clientInstance *mongo.Client
var databaseInstance *mongo.Database

func Connect() (*mongo.Database, error) {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	var clientErr error

	if clientInstance == nil {
		mutex.Lock()

		client, connErr := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

		if connErr != nil {
			clientErr = connErr
			return nil, connErr
		}

		log.Println("Initializing client")

		clientInstance = client
		mutex.Unlock()
	}

	database := os.Getenv("DATABASE_NAME")

	if database == "" {
		log.Fatal("You must set your 'DATABASE_NAME' environmental variable.")
		return nil, clientErr
	}

	if clientErr = clientInstance.Ping(context.TODO(), readpref.Primary()); clientErr != nil {
		log.Fatal(clientErr)
		return nil, clientErr
	}

	if databaseInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		databaseInstance = clientInstance.Database(database)
		log.Println("Initializing DB")
		log.Println("DB connected")
	}

	return databaseInstance, nil
}
