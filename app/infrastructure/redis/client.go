package redis

import (
	"time"

	"github.com/redis/go-redis/v9"

	"github/code-kakitai/code-kakitai/config"
)

var (
	redisClient *redis.Client
)

func GetRedisClient() *redis.Client {
	return redisClient
}

func SetRedisClient(c *redis.Client) {
	redisClient = c
}

func NewClient(conf config.Redis) *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:                  conf.Host + ":" + conf.Port,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return redisClient
}
