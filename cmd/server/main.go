package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Health": "The API is healthy!",
		})
	})

	router.Run(":8081")
}
