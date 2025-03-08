package log

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jomla97/loggernaut-api/database"
)

// Get handles the GET /log endpoint
func Get(c *gin.Context) {
	// Get the system and tags from the request headers
	t := c.Request.Header.Get("X-Meta-Tags")
	system := c.Param("system")
	id := strings.Trim(c.Param("id"), "/")

	// Find a single entry if an ID is provided
	if id != "" {
		fmt.Println("Finding entry with ID:", id)
		// Find the entry in the database
		entry, err := database.FindOne(system, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the entry
		c.JSON(http.StatusOK, entry)
		return
	}

	fmt.Println("Finding entries with tags:", t)

	// Split the tags into a slice
	var tags []string
	if t != "" {
		tags = strings.Split(t, ",")
	}

	// Find the entries in the database
	entries, err := database.Find(system, tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set entries to an empty slice if it is nil
	if entries == nil {
		entries = []interface{}{}
	}

	// Return the entries
	c.JSON(http.StatusOK, entries)
}
