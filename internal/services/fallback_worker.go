package services

import (
	"log"
	"encoding/json"

	"log-collector/internal/models"
	"log-collector/internal/config"
	
)

func StartFallbackWorker() {
	key := config.GetRedisFallbackKey()

	go func() {
		log.Println("[FallbackWorker] started, listening Redis fallback queue...")
		for {
			result, err := RedisClient.BLPop(ctx, 0, key).Result()
			if err != nil {
				log.Println("[FallbackWorker] BLPOP error:", err)
				continue
			}

			if len(result) < 2 {
				continue
			}

			rawLog := result[1]

			var entry models.LogEntry
			err = json.Unmarshal([]byte(rawLog), &entry)
			if err != nil {
				log.Println("[FallbackWorker] Unmarshal error:", err)
				continue
			}

			log.Printf("[FallbackWorker] Reprocessing log from service: %s", entry.Service)

			DB.Create(&entry)
			SendToSlack(entry)
			SendToElasticsearch(entry)
			SendMail(entry)
		}
	}()
}
