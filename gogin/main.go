package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Get the port from the PORT environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(":" + port)
}
