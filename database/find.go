package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// FindOne finds a single entry in the database by its ID
func FindOne(collection, idHex string) (entry interface{}, err error) {
	// Convert the ID to an ObjectID
	id, err := bson.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, err
	}

	// Find the entry in the database
	result := Client.Database(name).Collection(collection).FindOne(context.TODO(), bson.M{"_id": id})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	} else if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&entry); err != nil {
		return nil, err
	}

	return entry, nil
}

// Find finds entries in the database by their tags
func Find(collection string, tags []string) (entries []interface{}, err error) {
	filter := bson.M{}
	if len(tags) > 0 {
		filter["loggernaut_meta.source.tags"] = bson.M{
			"$all": tags,
		}
	}
	curs, err := Client.Database(name).Collection(collection).Find(context.TODO(), filter)
	if err != nil {
		return entries, err
	}

	for curs.Next(context.Background()) {
		var entry interface{}
		if err := curs.Decode(&entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
