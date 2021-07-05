package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dargoyaki/darbotyaki-backend/common"
)

var Router *gin.Engine

func main() {
	r := gin.Default()

	api := r.Group("/api")
	common.CommonsRegister(api.Group("/"))

	r.Run() // Runs on 0.0.0.0:8080
}
