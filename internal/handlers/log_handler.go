package handlers

import (
	"log-collector/internal/models"
	"log-collector/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLog(c *gin.Context) {
	var entry models.LogEntry

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid log format"})
		return
	}

	services.ProcessLog(entry)
	c.JSON(http.StatusOK, gin.H{"status": "log received"})
}

func GetLogs(c *gin.Context) {
	level := c.Query("level")
	service := c.Query("service")

	filteredLogs := services.FilterLogs(level, service)
	c.JSON(http.StatusOK, filteredLogs)
}



