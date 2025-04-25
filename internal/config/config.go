package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func GetSlackLogLevel() string {
	level := os.Getenv("SLACK_LOG_LEVEL")
	if level == "" {
		return "info" 
	}
	return strings.ToLower(level)
}


func IsElasticEnabled() bool {
	return os.Getenv("ELASTIC_ENABLED") == "true"
}

func GetElasticHost() string {
	return os.Getenv("ELASTIC_HOST")
}


func GetElasticIndex() string {
	index := os.Getenv("ELASTIC_INDEX")
	if index == "" {
		return "log-entries"
	}
	return index
}

func GetRedisHost() string {
	return os.Getenv("REDIS_HOST")
}

func GetRedisPort() string {
	return os.Getenv("REDIS_PORT")
}

func GetRedisFallbackKey() string {
	key := os.Getenv("REDIS_FALLBACK_KEY")
	if key == "" {
		return "fallback_logs"
	}
	return key
}

func GetLogChannelCapacity() int {
	val := os.Getenv("LOG_CHANNEL_CAPACITY")
	n, err := strconv.Atoi(val)
	if err != nil || n <= 0 {
		return 100
	}
	return n
}

func IsEmailEnabled() bool {
	return os.Getenv("EMAIL_ENABLED") == "true"
}

func GetMailLogLevel() string {
	level := os.Getenv("MAIL_LOG_LEVEL")
	if level == "" {
		return "error" 
	}
	return strings.ToLower(level)
}
