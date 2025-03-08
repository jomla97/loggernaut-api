package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client is the MongoDB client used to interact with the database
var Client *mongo.Client

// databaseName is the name of the database in which logs will be stored
const databaseName = "logs"

// Init initializes the database connection
func Init() {
	c, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	Client = c
}

// Close closes the database connection
func Close() {
	if err := Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
