package server

import "github.com/gin-gonic/gin"

// Start starts the web server
func Start() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/ingest", ingest)
	r.Run(":80")
}
