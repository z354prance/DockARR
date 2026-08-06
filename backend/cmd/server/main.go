package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const Version = "0.1.0-dev"

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": Version,
		})
	})

	router.Run(":8080")
}
