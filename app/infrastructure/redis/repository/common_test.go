package repository

import (
	"testing"

	redis "github.com/redis/go-redis/v9"

	infraRedis "github/code-kakitai/code-kakitai/infrastructure/redis"
)

var (
	redisCli *redis.Client
)

func TestMain(m *testing.M) {
	// テスト用Redisのセットアップ
	redisCli = infraRedis.NewTestClient()
	defer redisCli.Close()

	// テスト実行
	m.Run()
}
