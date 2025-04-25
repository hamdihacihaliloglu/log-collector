package services

import (
	"log-collector/internal/models"
)

var LogJobs chan models.LogEntry

func InitQueue(bufferSize int) {
	LogJobs = make(chan models.LogEntry, bufferSize)
}
