package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonsRegister(router *gin.RouterGroup) {
	router.GET("/", CommonsPing)
}

func CommonsPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello World!",
	})
}
