package redis

import (
	"github/code-kakitai/code-kakitai/config"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
)

func NewClient(conf *config.Config) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:                  conf.Redis.Addr,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return Client
}

func NewTestClient(conf *config.Config) *redis.Client {
	s, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr:                  s.Addr(),
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return client
}
