package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	serviceId := os.Getenv("ID")

	fmt.Printf("API Service #%s\n", serviceId)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": fmt.Sprintf("OK #%s", serviceId),
		})
	})
	r.Run("0.0.0.0:3000")
}
