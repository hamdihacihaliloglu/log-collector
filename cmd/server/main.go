package main

import (
	"log"
	"log-collector/internal/handlers"
	"log-collector/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	services.InitDB()
	services.InitQueue(100)    
	services.StartWorkers(5)    
	r := gin.Default()
	r.POST("/logs", handlers.PostLog)
	r.GET("/logs", handlers.GetLogs)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Log Collector started on :8080")
	r.Run(":8080")
}
