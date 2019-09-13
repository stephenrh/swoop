package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	app.Run(fmt.Sprintf(":%s", port))
}

func mainHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"info":    "etc",
	})
}
