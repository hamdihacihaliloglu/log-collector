package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log-collector/internal/config"
	"log-collector/internal/models"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.GetRedisHost(), config.GetRedisPort()),
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("[Redis] Connection error:", err)
	}
	log.Println("[Redis] Connected")
}

func PushToRedisFallback(entry models.LogEntry) {
	key := config.GetRedisFallbackKey()

	data, err := json.Marshal(entry)
	if err != nil {
		log.Println("[Redis] Marshal error:", err)
		return
	}

	err = RedisClient.RPush(ctx, key, data).Err()
	if err != nil {
		log.Println("[Redis] RPUSH error:", err)
		return
	}

	log.Println("[Redis] Pushed log to fallback queue:", key)
}
