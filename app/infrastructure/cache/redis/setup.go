package redis

import (
	"github/code-kakitai/code-kakitai/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
)

func Setup() func() {
	Client = redis.NewClient(&redis.Options{
		Addr:                  config.Config.ReidsConfig.Addr,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
	})

	return func() {
		_ = Client.Close()
	}
}
