package parsing

import (
	"fmt"

	"github.com/jomla97/loggernaut-api/database"
	"github.com/jomla97/loggernaut-api/inbox"
)

// shouldRun is a flag that indicates whether the parser should run, allowing it
// to start again after it has finished if more log files are added to the inbox
var shouldRun bool

// running is a flag that indicates whether the parser is currently running,
// preventing multiple instances from running concurrently
var running bool

// Start starts the parser, or if it's already running, signals it to start again
// after it has finished processing the current log files in the inbox
func Start() {
	shouldRun = true
	safeStart()
}

// safeStart starts a new parser goroutine if one is not already running,
// parsing all log and meta files in the inbox folder and inserting them into the database
func safeStart() {
	// Check if the parser is already running
	if running || !shouldRun {
		return
	}

	running = true
	shouldRun = false

	// Start the parser
	go func() {
		fmt.Println("Parsing started...")
		// Defer setting the parsing flag to false
		defer func() {
			running = false
		}()

		// Start the parser again if required
		defer safeStart()

		// Get all log files in the inbox
		logs, err := inbox.GetAll()
		if err != nil {
			fmt.Printf("Error: failed to get log files: %s\n", err.Error())
			return
		}

		// Parse each log file
		//TODO: concurrency
		for _, log := range logs {
			processLog(log)
		}
	}()
}

// processLog parses the specified log file, inserts it into the database and
// deletes the log file and its meta data file from the inbox
func processLog(log inbox.Log) {
	// Parse the log file
	entries, err := parse(log)
	if err != nil {
		fmt.Printf("Error: failed to parse log '%s': %s\n", log.ID, err.Error())
	}

	// Log a warning if no entries were parsed
	if len(entries) == 0 {
		fmt.Printf("Warning: no entries were parsed from log %s\n", log.ID)
		return
	}

	// Insert the parsed entries into the database
	err = database.Insert(log, entries)
	if err != nil {
		fmt.Printf("Error: failed to insert entries into database: %s\n", err.Error())
	}

	// Delete the log file and its meta data file
	err = log.Delete()
	if err != nil {
		fmt.Printf("failed to delete log file: %s\n", err.Error())
	}
}
