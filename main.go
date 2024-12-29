package main

import (
	"net/http"
	"plane/core"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	core.Initialize()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
