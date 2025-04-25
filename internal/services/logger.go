package services

import (
	"log"
	"log-collector/internal/config"
	"log-collector/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := config.GetDBConnectionString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.LogEntry{})

	DB = db
}

func ProcessLog(entry models.LogEntry) {
	select {
	case LogJobs <- entry:
		// success
	default:
		PushToRedisFallback(entry) // fallback if channel full
	}
}


func GetAllLogs() []models.LogEntry {
	var logs []models.LogEntry
	DB.Find(&logs)
	return logs
}

func FilterLogs(level, service string) []models.LogEntry {
	var logs []models.LogEntry
	query := DB.Model(&models.LogEntry{})

	if level != "" {
		query = query.Where("level = ?", level)
	}
	if service != "" {
		query = query.Where("service = ?", service)
	}

	query.Find(&logs)
	return logs
}
