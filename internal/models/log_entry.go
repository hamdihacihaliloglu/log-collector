package models

import "time"

type LogEntry struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Service   string    `json:"service"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp string    `json:"timestamp"` 
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
