package services

import (
	"bytes"
	"encoding/json"
	"log"
	"log-collector/internal/config"
	"log-collector/internal/models"
	"log-collector/internal/logging"
	"net/http"
)

func SendToSlack(entry models.LogEntry) {
	if !config.IsSlackEnabled() {
		log.Println("[Slack] Disabled via config")
		return
	}

	minLevel := config.GetSlackLogLevel()
	if !logging.ShouldLog(entry.Level, minLevel) {
		log.Printf("[Slack] Skipped %s log below threshold: %s\n", entry.Service, entry.Level)
		return
	}

	webhook := config.GetSlackWebhookURL()

	color := getColor(entry.Level)

	slackPayload := map[string]interface{}{
		"attachments": []map[string]interface{}{
			{
				"color": color,
				"fields": []map[string]string{
					{"title": "Service", "value": entry.Service, "short": "true"},
					{"title": "Level", "value": entry.Level, "short": "true"},
					{"title": "Message", "value": entry.Message},
					{"title": "Timestamp", "value": entry.Timestamp},
				},
			},
		},
	}

	jsonData, err := json.Marshal(slackPayload)
	if err != nil {
		log.Println("[Slack] JSON error:", err)
		return
	}

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("[Slack] Send error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[Slack] Status code: %d\n", resp.StatusCode)
	} else {
		log.Println("[Slack] Mesaj başarıyla gönderildi.")
	}
}

func getColor(level string) string {
	switch level {
	case "error":
		return "#FF0000"
	case "warning":
		return "#FFA500"
	case "info":
		return "#36a64f"
	default:
		return "#CCCCCC"
	}
}
