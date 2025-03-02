package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func postLog(c *gin.Context) {
	fileSystemPath := c.GetHeader("X-File-System-Path")
	system := c.GetHeader("X-System")
	tags := strings.Split(c.GetHeader("X-Tags"), ",")

	// Parse the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to parse form: %s", err.Error())})
		return
	}

	// Get the file header
	fileHeader, ok := form.File["log"]
	if !ok || len(fileHeader) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no log file provided"})
		return
	}

	// Open the file
	file, err := fileHeader[0].Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to open file: %s", err.Error())})
		return
	}

	// Read the file into a buffer
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to read file: %s", err.Error())})
		return
	}

	client.Database("logs").Collection("logs").InsertOne(context.TODO(), gin.H{
		"fileSystemPath": fileSystemPath,
		"system":         system,
		"tags":           tags,
		"log":            buf.String(),
	})

	c.Status(http.StatusCreated)
}
