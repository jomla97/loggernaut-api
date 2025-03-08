package database

import (
	"context"

	"github.com/jomla97/loggernaut-api/inbox"
)

// Insert the parsed entries of the specified log into the database
func Insert(log inbox.Log, entries []map[string]interface{}) error {
	for i, _ := range entries {
		entries[i]["loggernaut_meta"] = log.Meta
	}
	_, err := Client.
		Database(name).
		Collection(log.Meta.Source.System).
		InsertMany(context.TODO(), entries)
	return err
}
