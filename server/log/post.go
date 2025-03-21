package log

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jomla97/loggernaut-api/inbox"
	"github.com/jomla97/loggernaut-api/parsing"
)

// Post handles the POST /log endpoint
func Post(c *gin.Context) {
	// Parse the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to parse form: %s", err.Error())})
		return
	}

	// Get the file headers
	var logFileHeaders, metaFileHeaders []*multipart.FileHeader
	if _, ok := form.File["log"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no log files provided"})
		return
	} else {
		logFileHeaders = form.File["log"]
	}
	if _, ok := form.File["meta"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no meta files provided"})
		return
	} else {
		metaFileHeaders = form.File["meta"]
	}

	// Validate the file headers
	if len(logFileHeaders) != len(metaFileHeaders) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "number of log and meta files do not match"})
		return
	} else if len(logFileHeaders) == 0 || len(metaFileHeaders) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no log files provided"})
		return
	} else if len(form.File) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only log and meta files are allowed"})
		return
	}

	// Count the number of valid file pairs
	var validPairs int
	for _, fh := range logFileHeaders {
		for _, mfh := range metaFileHeaders {
			if fh.Filename == strings.TrimSuffix(mfh.Filename, ".meta.json") {
				validPairs++
				break
			}
		}
	}

	// Make sure the log and meta files match
	if validPairs != len(logFileHeaders) {
		fmt.Println("validPairs", validPairs)
		fmt.Println("len(logFileHeaders)", len(logFileHeaders))
		c.JSON(http.StatusBadRequest, gin.H{"error": "log and meta files do not match"})
		return
	}

	// Save the files to the inbox for later processing
	for _, fileHeader := range append(logFileHeaders, metaFileHeaders...) {
		err = c.SaveUploadedFile(fileHeader, filepath.Join(inbox.Path, fileHeader.Filename))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to save file: %s", err.Error())})
			return
		}
	}

	// Start the parsing process
	parsing.Start()

	// Return a success response
	c.Status(http.StatusCreated)
}
