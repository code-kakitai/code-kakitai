//go:build integration_write

package api_write_test

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/gin-gonic/gin"
	"gopkg.in/testfixtures.v2"

	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	dbTest "github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	infraRedis "github/code-kakitai/code-kakitai/infrastructure/redis"
	"github/code-kakitai/code-kakitai/presentation/settings"
	"github/code-kakitai/code-kakitai/server/route"

	redis "github.com/redis/go-redis/v9"
)

var (
	fixtures *testfixtures.Context
	api      *gin.Engine
)

func TestMain(m *testing.M) {
	var err error

	// DBの立ち上げ
	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	// DBへ接続する
	dbCon := dbTest.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBをセットアップ
	dbTest.SetupTestDB("../../../infrastructure/mysql/db/schema/schema.sql")

	// テストデータの準備
	fixtures, err = testfixtures.NewFolder(
		dbCon,
		&testfixtures.MySQL{},
		"../../../infrastructure/mysql/fixtures",
	)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetQuery(q)
	db.SetReadQuery(q)
	db.SetDB(dbCon)
	db.SetReadDB(dbCon)

	infraRedis.SetRedisClient(NewTestClient())
	cli := infraRedis.GetRedisClient()
	defer cli.Close()

	api = settings.NewGinEngine()
	route.InitRoute(api)

	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
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
