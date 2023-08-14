package query_service

import (
	infraRedis "github/code-kakitai/code-kakitai/infrastructure/cache/redis"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"testing"

	redis "github.com/redis/go-redis/v9"
)

var (
	query    *dbgen.Queries
	redisCli *redis.Client
)

func TestMain(m *testing.M) {
	// DBの立ち上げ
	resource, pool := db_test.CreateContainer()
	defer db_test.CloseContainer(resource, pool)

	// DBへ接続する
	db := db_test.ConnectDB(resource, pool)
	defer db.Close()

	// テスト用DBをセットアップ
	db_test.SetupTestDB()

	query = dbgen.New(db)

	// テスト用Redisのセットアップ
	redisCli = infraRedis.NewTestClient()
	defer redisCli.Close()

	// テスト実行
	m.Run()
}
