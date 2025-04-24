package config

import (
	"fmt"
	"os"
)

func GetDBConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
}

func IsSlackEnabled() bool {
	return os.Getenv("SLACK_ENABLED") == "true"
}

func GetSlackWebhookURL() string {
	return os.Getenv("SLACK_WEBHOOK_URL")
}
