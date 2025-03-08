package database

import (
	"context"

	"github.com/jomla97/loggernaut-api/inbox"
)

// Insert the parsed entries of the specified log into the database
func Insert(log inbox.Log, entries []interface{}) error {
	//TODO: insert meta data
	//TODO: fix, currently it does not work
	_, err := Client.Database(databaseName).Collection(log.Meta.Source.System).InsertMany(context.TODO(), entries)
	return err
}
