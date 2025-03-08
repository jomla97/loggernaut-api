package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Client is the MongoDB client used to interact with the database
var Client *mongo.Client

// name is the name of the database in which logs will be stored
const name = "logs"

// uri is the URI of the MongoDB instance
const uri = "mongodb://loggernaut-db:27017/loggernaut"

// Init initializes the database connection
func Init() {
	c, err := mongo.Connect(options.Client().ApplyURI(uri))
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
