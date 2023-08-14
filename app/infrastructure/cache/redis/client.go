package redis

import (
	"github/code-kakitai/code-kakitai/config"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
)

var (
	redisCli *redis.Client
)

func NewClient(conf *config.Config) *redis.Client {
	redisCli = redis.NewClient(&redis.Options{
		Addr:                  conf.Redis.Addr,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return redisCli
}

func NewTestClient() *redis.Client {
	s, _ := miniredis.Run()
	redisCli := redis.NewClient(&redis.Options{
		Addr:                  s.Addr(),
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return redisCli
}
