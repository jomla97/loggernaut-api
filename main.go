package main

import (
	"github.com/jomla97/loggernaut-api/database"
	"github.com/jomla97/loggernaut-api/parsing"
	"github.com/jomla97/loggernaut-api/server"
)

func main() {
	// Initialize the database into which logs will be inserted
	database.Init()
	defer database.Close()

	// Start the parser, which will be restarted when needed as logs are received
	parsing.Start()

	// Start the server, which will listen for incoming requests
	server.Start()
}
