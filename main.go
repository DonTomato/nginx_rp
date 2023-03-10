package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const version = 3

func main() {
	serviceId := os.Getenv("ID")
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	fmt.Printf("API Service v%d #%s\n", version, serviceId)

	if name != "" {
		fmt.Printf("Hello %s\n", name)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": fmt.Sprintf("OK v%d: #%s", version, serviceId),
		})
	})
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
