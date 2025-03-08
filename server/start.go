package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jomla97/loggernaut-api/server/log"
)

// Start starts the web server
func Start() {
	//TODO: oauth
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/log/:system/*id", log.Get)
	r.POST("/log", log.Post)
	r.Run(":80")
}
