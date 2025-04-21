package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
