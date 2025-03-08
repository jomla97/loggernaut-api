package parsing

import (
	"fmt"
	"reflect"

	"github.com/jomla97/loggernaut-api/inbox"
)

// Parser is an interface for log parsers, which can parse log data into a list of entries
type Parser interface {
	Parse(data *[]byte) (entries []interface{}, err error)
}

// parse parses the log at the specified path and inserts it into the database
func parse(log inbox.Log) (entries []interface{}, err error) {
	// Read the log file data
	data, err := log.Read()
	if err != nil {
		return entries, fmt.Errorf("failed to read log file: %w", err)
	}

	// Create a list of parsers
	parsers := []Parser{JSONParser{}, GrokParser{}}

	// Iterate over each parser and attempt to parse the log data
	for _, parser := range parsers {
		// Attempt to parse the log data
		entries, err := parser.Parse(&data)
		if err != nil {
			continue
		}

		// Log information about the parsing process
		fmt.Printf(
			"Parsed %d entries from log %s using parser %s\n",
			len(entries), log.ID, reflect.TypeOf(parser).Name(),
		)

		return entries, nil
	}

	// Return an error if no parser was able to parse the log data
	return entries, fmt.Errorf("no parser was able to parse log %s", log.ID)
}
