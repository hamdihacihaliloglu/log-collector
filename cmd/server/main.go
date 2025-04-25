package main

import (
	"log"
	"log-collector/internal/handlers"
	"log-collector/internal/services"
	"log-collector/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	services.InitDB()
	services.InitRedis()
	services.InitQueue(config.GetLogChannelCapacity())
	services.StartWorkers(5) 
	services.StartFallbackWorker() 
	 
	r := gin.Default()
	r.POST("/logs", handlers.PostLog)
	r.GET("/logs", handlers.GetLogs)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Log Collector started on :8080")
	r.Run(":8080")
}
