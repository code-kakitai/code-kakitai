package redis

import (
	"time"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"

	"github/code-kakitai/code-kakitai/config"
)

var (
	redisClient *redis.Client
)

func GetRedisClient() *redis.Client {
	return redisClient
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

func NewTestClient() *redis.Client {
	s, _ := miniredis.Run()
	redisClient := redis.NewClient(&redis.Options{
		Addr:                  s.Addr(),
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return redisClient
}
