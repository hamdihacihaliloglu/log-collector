package services

import (
	"bytes"
	"encoding/json"
	"log"
	"log-collector/internal/config"
	"log-collector/internal/models"
	"net/http"
)

func SendToElasticsearch(entry models.LogEntry) {
	if !config.IsElasticEnabled() {
		log.Println("[Elastic] Disabled via config")
		return
	}

	url := config.GetElasticHost() + "/" + config.GetElasticIndex() + "/_doc"

	data, err := json.Marshal(entry)
	if err != nil {
		log.Println("[Elastic] JSON marshal error:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("[Elastic] Post error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		log.Printf("[Elastic] Unexpected status code: %d\n", resp.StatusCode)
	} else {
		log.Println("[Elastic] Log sent to Elasticsearch.")
	}
}
