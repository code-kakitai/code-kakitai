package redis

import (
	"github/code-kakitai/code-kakitai/config"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
)

var (
	Cli *redis.Client
)

func NewClient(conf config.Redis) *redis.Client {
	Cli = redis.NewClient(&redis.Options{
		Addr:                  conf.Host + ":" + conf.Port,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return Cli
}

func NewTestClient() *redis.Client {
	s, _ := miniredis.Run()
	Cli := redis.NewClient(&redis.Options{
		Addr:                  s.Addr(),
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return Cli
}
