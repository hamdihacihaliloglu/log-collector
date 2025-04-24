package main

import (
	"log"
	"log-collector/internal/handlers"
	"log-collector/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	services.InitDB()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/logs", handlers.PostLog)

	r.GET("/logs", handlers.GetLogs)

	log.Println("Starting server on :8080")
	r.Run(":8080")
}
