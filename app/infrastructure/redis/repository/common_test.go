package repository

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	redis "github.com/redis/go-redis/v9"

	infraRedis "github/code-kakitai/code-kakitai/infrastructure/redis"
)

func TestMain(m *testing.M) {
	// テスト用Redisのセットアップ
	infraRedis.SetRedisClient(NewTestClient())
	cli := infraRedis.GetRedisClient()
	defer cli.Close()

	// テスト実行
	m.Run()
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
